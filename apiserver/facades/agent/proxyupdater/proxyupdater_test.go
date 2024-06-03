// Copyright 2016 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package proxyupdater_test

import (
	"context"
	"time"

	"github.com/juju/names/v5"
	"github.com/juju/testing"
	jc "github.com/juju/testing/checkers"
	"github.com/juju/worker/v4/workertest"
	"go.uber.org/mock/gomock"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/apiserver/common"
	"github.com/juju/juju/apiserver/facades/agent/proxyupdater"
	apiservertesting "github.com/juju/juju/apiserver/testing"
	"github.com/juju/juju/controller"
	"github.com/juju/juju/core/network"
	"github.com/juju/juju/core/watcher/watchertest"
	"github.com/juju/juju/rpc/params"
	"github.com/juju/juju/state"
	coretesting "github.com/juju/juju/testing"
)

type ProxyUpdaterSuite struct {
	coretesting.BaseSuite

	state      *stubBackend
	resources  *common.Resources
	authorizer apiservertesting.FakeAuthorizer
	facade     *proxyupdater.API
	tag        names.MachineTag

	modelConfigService      *MockModelConfigService
	controllerConfigService *MockControllerConfigService
}

var _ = gc.Suite(&ProxyUpdaterSuite{})

func (s *ProxyUpdaterSuite) SetUpSuite(c *gc.C) {
	s.BaseSuite.SetUpSuite(c)
}

func (s *ProxyUpdaterSuite) SetUpTest(c *gc.C) {
	s.BaseSuite.SetUpTest(c)
	s.resources = common.NewResources()
	s.AddCleanup(func(_ *gc.C) { s.resources.StopAll() })
	s.authorizer = apiservertesting.FakeAuthorizer{
		Tag:        names.NewMachineTag("1"),
		Controller: false,
	}
	s.tag = names.NewMachineTag("1")
	s.state = &stubBackend{}
	s.state.SetUp(c)
	s.AddCleanup(func(_ *gc.C) { s.state.Kill() })
}

func (s *ProxyUpdaterSuite) setupAPI(c *gc.C) *gomock.Controller {
	ctrl := gomock.NewController(c)

	s.controllerConfigService = NewMockControllerConfigService(ctrl)
	s.modelConfigService = NewMockModelConfigService(ctrl)

	api, err := proxyupdater.NewAPIV2(s.state, s.controllerConfigService, s.modelConfigService, s.resources, s.authorizer)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(api, gc.NotNil)
	s.facade = api

	// Shouldn't have any calls yet
	apiservertesting.CheckMethodCalls(c, s.state.Stub)

	return ctrl
}

func (s *ProxyUpdaterSuite) TestWatchForProxyConfigAndAPIHostPortChanges(c *gc.C) {
	ctrl := s.setupAPI(c)
	defer ctrl.Finish()

	// WatchForProxyConfigAndAPIHostPortChanges combines WatchForModelConfigChanges
	// and WatchAPIHostPorts. Check that they are both called.

	// Create fake model config watcher preloaded with one item in the channel
	modelConfigChanges := make(chan []string, 1)
	modelConfigChanges <- []string{}
	modelConfigWatcher := watchertest.NewMockStringsWatcher(modelConfigChanges)
	defer modelConfigWatcher.Kill()
	s.modelConfigService.EXPECT().Watch().Return(modelConfigWatcher, nil)

	result := s.facade.WatchForProxyConfigAndAPIHostPortChanges(context.Background(), s.oneEntity())
	c.Assert(result.Results, gc.HasLen, 1)
	c.Assert(result.Results[0].Error, gc.IsNil)

	s.state.Stub.CheckCallNames(c,
		"WatchAPIHostPortsForAgents",
	)

	// Verify the watcher resource was registered.
	c.Assert(s.resources.Count(), gc.Equals, 1)
	resource := s.resources.Get(result.Results[0].NotifyWatcherId)
	watcher, ok := resource.(state.NotifyWatcher)
	c.Assert(ok, jc.IsTrue)

	// Verify the initial event was consumed.
	select {
	case <-watcher.Changes():
		c.Fatalf("initial event never consumed")
	case <-time.After(coretesting.ShortWait):
	}
}

func (s *ProxyUpdaterSuite) oneEntity() params.Entities {
	entities := params.Entities{
		make([]params.Entity, 1),
	}
	entities.Entities[0].Tag = s.tag.String()
	return entities
}

func (s *ProxyUpdaterSuite) TestMirrorConfig(c *gc.C) {
	ctrl := s.setupAPI(c)
	defer ctrl.Finish()

	// Check that the ProxyConfig combines data from ModelConfig and APIHostPorts
	s.modelConfigService.EXPECT().ModelConfig(gomock.Any()).Return(coretesting.CustomModelConfig(c,
		coretesting.Attrs{
			"apt-mirror": "http://mirror",
		},
	), nil)
	s.controllerConfigService.EXPECT().ControllerConfig(gomock.Any()).Return(coretesting.FakeControllerConfig(), nil)

	cfg := s.facade.ProxyConfig(context.Background(), s.oneEntity())

	s.state.Stub.CheckCallNames(c,
		"APIHostPortsForAgents",
	)

	c.Assert(cfg.Results, gc.HasLen, 1)
	c.Assert(cfg.Results[0].AptMirror, gc.Equals, "http://mirror")
}

func (s *ProxyUpdaterSuite) TestProxyConfig(c *gc.C) {
	ctrl := s.setupAPI(c)
	defer ctrl.Finish()

	// Check that the ProxyConfig combines data from ModelConfig and APIHostPorts
	s.modelConfigService.EXPECT().ModelConfig(gomock.Any()).Return(coretesting.CustomModelConfig(c,
		coretesting.Attrs{
			"http-proxy":      "http proxy",
			"https-proxy":     "https proxy",
			"apt-http-proxy":  "apt http proxy",
			"apt-https-proxy": "apt https proxy",
		},
	), nil)
	s.controllerConfigService.EXPECT().ControllerConfig(gomock.Any()).Return(coretesting.FakeControllerConfig(), nil)

	cfg := s.facade.ProxyConfig(context.Background(), s.oneEntity())
	s.state.Stub.CheckCallNames(c,
		"APIHostPortsForAgents",
	)

	expectedLegacyNoProxy := "0.1.2.3,0.1.2.4,0.1.2.5"
	expectedJujuNoProxy := ""

	r := params.ProxyConfigResult{
		LegacyProxySettings: params.ProxyConfig{
			HTTP: "http proxy", HTTPS: "https proxy", FTP: "", NoProxy: expectedLegacyNoProxy},
		JujuProxySettings: params.ProxyConfig{
			HTTP: "", HTTPS: "", FTP: "", NoProxy: expectedJujuNoProxy},
		APTProxySettings: params.ProxyConfig{
			HTTP: "http://apt http proxy", HTTPS: "https://apt https proxy", FTP: "", NoProxy: ""},
	}
	c.Assert(cfg.Results[0], jc.DeepEquals, r)
}

func (s *ProxyUpdaterSuite) TestProxyConfigJujuProxy(c *gc.C) {
	ctrl := s.setupAPI(c)
	defer ctrl.Finish()

	s.modelConfigService.EXPECT().ModelConfig(gomock.Any()).Return(coretesting.CustomModelConfig(c,
		coretesting.Attrs{
			"juju-http-proxy":  "http proxy",
			"juju-https-proxy": "https proxy",
			"apt-http-proxy":   "apt http proxy",
			"apt-https-proxy":  "apt https proxy",
		},
	), nil)
	s.controllerConfigService.EXPECT().ControllerConfig(gomock.Any()).Return(coretesting.FakeControllerConfig(), nil)

	cfg := s.facade.ProxyConfig(context.Background(), s.oneEntity())
	s.state.Stub.CheckCallNames(c,
		"APIHostPortsForAgents",
	)

	// need to make sure that auto-population/auto-appending of controller IPs to
	// no-proxy is aware of which proxy settings are used: if non-legacy ones are used
	// then juju-no-proxy should be auto-modified
	expectedJujuNoProxy := "0.1.2.3,0.1.2.4,0.1.2.5"
	expectedLegacyNoProxy := ""

	r := params.ProxyConfigResult{
		JujuProxySettings: params.ProxyConfig{
			HTTP: "http proxy", HTTPS: "https proxy", FTP: "", NoProxy: expectedJujuNoProxy},
		LegacyProxySettings: params.ProxyConfig{
			HTTP: "", HTTPS: "", FTP: "", NoProxy: expectedLegacyNoProxy},
		APTProxySettings: params.ProxyConfig{
			HTTP: "http://apt http proxy", HTTPS: "https://apt https proxy", FTP: "", NoProxy: ""},
	}
	c.Assert(cfg.Results[0], jc.DeepEquals, r)
}

func (s *ProxyUpdaterSuite) TestProxyConfigExtendsExisting(c *gc.C) {
	ctrl := s.setupAPI(c)
	defer ctrl.Finish()

	// Check that the ProxyConfig combines data from ModelConfig and APIHostPorts
	s.modelConfigService.EXPECT().ModelConfig(gomock.Any()).Return(coretesting.CustomModelConfig(c,
		coretesting.Attrs{
			"http-proxy":      "http proxy",
			"https-proxy":     "https proxy",
			"apt-http-proxy":  "apt http proxy",
			"apt-https-proxy": "apt https proxy",
			"no-proxy":        "9.9.9.9",
		},
	), nil)
	s.controllerConfigService.EXPECT().ControllerConfig(gomock.Any()).Return(coretesting.FakeControllerConfig(), nil)

	cfg := s.facade.ProxyConfig(context.Background(), s.oneEntity())
	s.state.Stub.CheckCallNames(c,
		"APIHostPortsForAgents",
	)

	expectedNoProxy := "0.1.2.3,0.1.2.4,0.1.2.5,9.9.9.9"
	expectedAptNoProxy := "9.9.9.9"

	c.Assert(cfg.Results[0], jc.DeepEquals, params.ProxyConfigResult{
		LegacyProxySettings: params.ProxyConfig{
			HTTP: "http proxy", HTTPS: "https proxy", FTP: "", NoProxy: expectedNoProxy},
		APTProxySettings: params.ProxyConfig{
			HTTP: "http://apt http proxy", HTTPS: "https://apt https proxy", FTP: "", NoProxy: expectedAptNoProxy},
	})
}

func (s *ProxyUpdaterSuite) TestProxyConfigNoDuplicates(c *gc.C) {
	ctrl := s.setupAPI(c)
	defer ctrl.Finish()

	// Check that the ProxyConfig combines data from ModelConfig and APIHostPorts
	s.modelConfigService.EXPECT().ModelConfig(gomock.Any()).Return(coretesting.CustomModelConfig(c,
		coretesting.Attrs{
			"http-proxy":      "http proxy",
			"https-proxy":     "https proxy",
			"apt-http-proxy":  "apt http proxy",
			"apt-https-proxy": "apt https proxy",
			"no-proxy":        "0.1.2.3",
		},
	), nil)
	s.controllerConfigService.EXPECT().ControllerConfig(gomock.Any()).Return(coretesting.FakeControllerConfig(), nil)

	cfg := s.facade.ProxyConfig(context.Background(), s.oneEntity())
	s.state.Stub.CheckCallNames(c,
		"APIHostPortsForAgents",
	)

	expectedNoProxy := "0.1.2.3,0.1.2.4,0.1.2.5"
	expectedAptNoProxy := "0.1.2.3"

	c.Assert(cfg.Results[0], jc.DeepEquals, params.ProxyConfigResult{
		LegacyProxySettings: params.ProxyConfig{
			HTTP: "http proxy", HTTPS: "https proxy", FTP: "", NoProxy: expectedNoProxy},
		APTProxySettings: params.ProxyConfig{
			HTTP: "http://apt http proxy", HTTPS: "https://apt https proxy", FTP: "", NoProxy: expectedAptNoProxy},
	})
}

func (s *ProxyUpdaterSuite) TestSnapProxyConfig(c *gc.C) {
	ctrl := s.setupAPI(c)
	defer ctrl.Finish()

	s.modelConfigService.EXPECT().ModelConfig(gomock.Any()).Return(coretesting.CustomModelConfig(c,
		coretesting.Attrs{
			"snap-http-proxy":       "http proxy",
			"snap-https-proxy":      "https proxy",
			"snap-store-proxy":      "store proxy",
			"snap-store-assertions": "trust us",
		},
	), nil)
	s.controllerConfigService.EXPECT().ControllerConfig(gomock.Any()).Return(coretesting.FakeControllerConfig(), nil)

	cfg := s.facade.ProxyConfig(context.Background(), s.oneEntity())
	s.state.Stub.CheckCallNames(c,
		"APIHostPortsForAgents",
	)

	expectedNoProxy := "0.1.2.3,0.1.2.4,0.1.2.5"

	c.Assert(cfg.Results[0], jc.DeepEquals, params.ProxyConfigResult{
		LegacyProxySettings: params.ProxyConfig{NoProxy: expectedNoProxy},
		SnapProxySettings: params.ProxyConfig{
			HTTP: "http proxy", HTTPS: "https proxy"},
		SnapStoreProxyId:         "store proxy",
		SnapStoreProxyAssertions: "trust us",
	})
}

type stubBackend struct {
	*testing.Stub
	c         *gc.C
	hpWatcher workertest.NotAWatcher
}

func (sb *stubBackend) SetUp(c *gc.C) {
	sb.Stub = &testing.Stub{}
	sb.c = c
	sb.hpWatcher = workertest.NewFakeWatcher(1, 1)
}

func (sb *stubBackend) Kill() {
	sb.hpWatcher.Kill()
}

func (sb *stubBackend) APIHostPortsForAgents(_ controller.Config) ([]network.SpaceHostPorts, error) {
	sb.MethodCall(sb, "APIHostPortsForAgents")
	if err := sb.NextErr(); err != nil {
		return nil, err
	}
	hps := []network.SpaceHostPorts{
		network.NewSpaceHostPorts(1234, "0.1.2.3"),
		network.NewSpaceHostPorts(1234, "0.1.2.4"),
		network.NewSpaceHostPorts(1234, "0.1.2.5"),
	}
	return hps, nil
}

func (sb *stubBackend) WatchAPIHostPortsForAgents() state.NotifyWatcher {
	sb.MethodCall(sb, "WatchAPIHostPortsForAgents")
	return sb.hpWatcher
}
