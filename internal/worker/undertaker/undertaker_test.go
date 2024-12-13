// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package undertaker_test

import (
	"context"
	"time"

	"github.com/juju/clock/testclock"
	"github.com/juju/errors"
	"github.com/juju/testing"
	jc "github.com/juju/testing/checkers"
	"github.com/juju/worker/v4"
	"github.com/juju/worker/v4/workertest"
	"go.uber.org/mock/gomock"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/core/life"
	"github.com/juju/juju/core/watcher"
	"github.com/juju/juju/core/watcher/watchertest"
	"github.com/juju/juju/environs"
	"github.com/juju/juju/environs/cloudspec"
	"github.com/juju/juju/environs/envcontext"
	loggertesting "github.com/juju/juju/internal/logger/testing"
	"github.com/juju/juju/internal/worker/undertaker"
	"github.com/juju/juju/rpc/params"
)

// OldUndertakerSuite is *not* complete. But it's a lot more so
// than it was before, and should be much easier to extend.
type OldUndertakerSuite struct {
	testing.IsolationSuite
	fix fixture
}

var _ = gc.Suite(&OldUndertakerSuite{})

func (s *OldUndertakerSuite) SetUpTest(c *gc.C) {
	s.IsolationSuite.SetUpTest(c)
	minute := time.Minute
	s.fix = fixture{
		clock: testclock.NewDilatedWallClock(10 * time.Millisecond),
		info: params.UndertakerModelInfoResult{
			Result: params.UndertakerModelInfo{
				Life:           "dying",
				DestroyTimeout: &minute,
			},
		},
	}
}

func (s *OldUndertakerSuite) TestAliveError(c *gc.C) {
	s.fix.info.Result.Life = "alive"
	s.fix.dirty = true
	stub := s.fix.run(c, func(w worker.Worker) {
		err := workertest.CheckKilled(c, w)
		c.Check(err, gc.ErrorMatches, "model still alive")
	})
	stub.CheckCallNames(c, "WatchModel", "ModelInfo")
}

func (s *OldUndertakerSuite) TestAlreadyDeadRemoves(c *gc.C) {
	s.fix.info.Result.Life = "dead"
	stub := s.fix.run(c, func(w worker.Worker) {
		workertest.CheckKilled(c, w)
	})
	stub.CheckCallNames(c, "WatchModel", "ModelInfo", "ModelConfig", "CloudSpec", "Destroy", "RemoveModel")
}

func (s *OldUndertakerSuite) TestDyingDeadRemoved(c *gc.C) {
	stub := s.fix.run(c, func(w worker.Worker) {
		workertest.CheckKilled(c, w)
	})
	stub.CheckCallNames(c,
		"WatchModel",
		"ModelInfo",
		"WatchModelResources",
		"ProcessDyingModel",
		"ModelConfig",
		"CloudSpec",
		"Destroy",
		"RemoveModel",
	)
}

func (s *OldUndertakerSuite) TestControllerStopsWhenModelDead(c *gc.C) {
	s.fix.info.Result.IsSystem = true
	stub := s.fix.run(c, func(w worker.Worker) {
		workertest.CheckKilled(c, w)
	})
	stub.CheckCallNames(c,
		"WatchModel",
		"ModelInfo",

		"WatchModelResources",
		"ProcessDyingModel",
	)
}

func (s *OldUndertakerSuite) TestModelInfoErrorFatal(c *gc.C) {
	s.fix.errors = []error{nil, errors.New("pow")}
	s.fix.dirty = true
	stub := s.fix.run(c, func(w worker.Worker) {
		err := workertest.CheckKilled(c, w)
		c.Check(err, gc.ErrorMatches, "pow")
	})
	stub.CheckCallNames(c, "WatchModel", "ModelInfo")
}

func (s *OldUndertakerSuite) TestWatchModelResourcesErrorFatal(c *gc.C) {
	s.fix.errors = []error{nil, nil, errors.New("pow")}
	s.fix.dirty = true
	stub := s.fix.run(c, func(w worker.Worker) {
		err := workertest.CheckKilled(c, w)
		c.Check(err, gc.ErrorMatches, "proccesing model death: pow")
	})
	stub.CheckCallNames(c, "WatchModel", "ModelInfo", "WatchModelResources")
}

func (s *OldUndertakerSuite) TestProcessDyingModelErrorRetried(c *gc.C) {
	s.fix.errors = []error{
		nil, // WatchModel
		nil, // ModelInfo
		nil, // WatchModelResources,
		&params.Error{Code: params.CodeHasHostedModels},
		&params.Error{Code: params.CodeModelNotEmpty},
		nil, // ProcessDyingModel,
		nil, // ModelConfig
		nil, // CloudSpec
		nil, // Destroy,
		nil, // RemoveModel
	}
	stub := s.fix.run(c, func(w worker.Worker) {
		workertest.CheckKilled(c, w)
	})
	stub.CheckCallNames(c,
		"WatchModel",
		"ModelInfo",
		"WatchModelResources",
		"ProcessDyingModel",
		"ProcessDyingModel",
		"ProcessDyingModel",
		"ModelConfig",
		"CloudSpec",
		"Destroy",
		"RemoveModel",
	)
}

func (s *OldUndertakerSuite) TestProcessDyingModelErrorFatal(c *gc.C) {
	s.fix.errors = []error{
		nil, // WatchModel
		nil, // ModelInfo
		nil, // WatchModelResources,
		errors.New("nope"),
	}
	s.fix.dirty = true
	stub := s.fix.run(c, func(w worker.Worker) {
		err := workertest.CheckKilled(c, w)
		c.Check(err, gc.ErrorMatches, "proccesing model death: nope")
	})
	stub.CheckCallNames(c,
		"WatchModel",
		"ModelInfo",
		"WatchModelResources",
		"ProcessDyingModel",
	)
}

func (s *OldUndertakerSuite) TestDestroyErrorFatal(c *gc.C) {
	s.fix.errors = []error{nil, nil, nil, nil, errors.New("pow")}
	s.fix.info.Result.Life = "dead"
	s.fix.dirty = true
	stub := s.fix.run(c, func(w worker.Worker) {
		err := workertest.CheckKilled(c, w)
		c.Check(err, gc.ErrorMatches, "cannot destroy cloud resources: process destroy environ: pow")
	})
	stub.CheckCallNames(c, "WatchModel", "ModelInfo", "ModelConfig", "CloudSpec", "Destroy")
}

func (s *OldUndertakerSuite) TestDestroyErrorForced(c *gc.C) {
	s.fix.errors = []error{nil, nil, nil, nil, errors.New("pow")}
	s.fix.info.Result.Life = "dead"
	s.fix.info.Result.ForceDestroyed = true
	destroyTimeout := 500 * time.Millisecond
	s.fix.info.Result.DestroyTimeout = &destroyTimeout
	stub := s.fix.run(c, func(w worker.Worker) {
		err := workertest.CheckKilled(c, w)
		c.Assert(err, jc.ErrorIsNil)
	})
	// Removal continues despite the error calling destroy.
	mainCalls, destroyCloudCalls := s.sortCalls(c, stub)
	c.Assert(mainCalls, jc.DeepEquals, []string{"WatchModel", "ModelInfo", "RemoveModel"})
	c.Assert(destroyCloudCalls, jc.DeepEquals, []string{"ModelConfig", "CloudSpec", "Destroy"})
}

func (s *OldUndertakerSuite) TestRemoveModelErrorFatal(c *gc.C) {
	s.fix.errors = []error{nil, nil, nil, nil, nil, errors.New("pow")}
	s.fix.info.Result.Life = "dead"
	s.fix.dirty = true
	stub := s.fix.run(c, func(w worker.Worker) {
		err := workertest.CheckKilled(c, w)
		c.Check(err, gc.ErrorMatches, "cannot remove model: pow")
	})
	mainCalls, destroyCloudCalls := s.sortCalls(c, stub)
	c.Assert(mainCalls, jc.DeepEquals, []string{"WatchModel", "ModelInfo", "RemoveModel"})
	c.Assert(destroyCloudCalls, jc.DeepEquals, []string{"ModelConfig", "CloudSpec", "Destroy"})
}

func (s *OldUndertakerSuite) TestDestroyTimeout(c *gc.C) {
	notEmptyErr := &params.Error{Code: params.CodeModelNotEmpty}
	s.fix.errors = []error{nil, nil, nil, notEmptyErr, notEmptyErr, notEmptyErr, notEmptyErr, errors.Timeoutf("error")}
	s.fix.dirty = true
	s.fix.advance = 2 * time.Minute
	stub := s.fix.run(c, func(w worker.Worker) {
		err := workertest.CheckKilled(c, w)
		c.Assert(err, gc.ErrorMatches, ".* timeout")
	})
	// Depending on timing there can be 1 or more ProcessDyingModel calls.
	calls := stub.Calls()
	var callNames []string
	for _, call := range calls {
		if call.FuncName == "ProcessDyingModel" {
			continue
		}
		callNames = append(callNames, call.FuncName)
	}
	c.Assert(callNames, jc.DeepEquals, []string{"WatchModel", "ModelInfo", "WatchModelResources"})
}

func (s *OldUndertakerSuite) TestDestroyTimeoutForce(c *gc.C) {
	s.fix.info.Result.ForceDestroyed = true
	s.fix.advance = 2 * time.Minute
	stub := s.fix.run(c, func(w worker.Worker) {
		err := workertest.CheckKilled(c, w)
		c.Assert(err, jc.ErrorIsNil)
	})
	mainCalls, destroyCloudCalls := s.sortCalls(c, stub)
	c.Assert(mainCalls, jc.DeepEquals, []string{"WatchModel", "ModelInfo", "WatchModelResources", "ProcessDyingModel", "RemoveModel"})
	c.Assert(destroyCloudCalls, jc.DeepEquals, []string{"ModelConfig", "CloudSpec", "Destroy"})
}

func (s *OldUndertakerSuite) TestEnvironDestroyTimeout(c *gc.C) {
	timeout := time.Millisecond
	s.fix.info.Result.DestroyTimeout = &timeout
	s.fix.dirty = true
	stub := s.fix.run(c, func(w worker.Worker) {
		err := workertest.CheckKilled(c, w)
		c.Assert(err, jc.ErrorIsNil)
	})
	mainCalls, destroyCloudCalls := s.sortCalls(c, stub)
	c.Assert(mainCalls, jc.DeepEquals, []string{"WatchModel", "ModelInfo", "WatchModelResources", "ProcessDyingModel", "RemoveModel"})
	c.Assert(destroyCloudCalls, jc.DeepEquals, []string{"ModelConfig", "CloudSpec", "Destroy"})
}

func (s *OldUndertakerSuite) TestEnvironDestroyTimeoutForce(c *gc.C) {
	timeout := time.Second
	s.fix.info.Result.DestroyTimeout = &timeout
	s.fix.info.Result.ForceDestroyed = true
	s.fix.dirty = true
	stub := s.fix.run(c, func(w worker.Worker) {
		err := workertest.CheckKilled(c, w)
		c.Assert(err, jc.ErrorIsNil)
	})
	mainCalls, destroyCloudCalls := s.sortCalls(c, stub)
	c.Assert(mainCalls, jc.DeepEquals, []string{"WatchModel", "ModelInfo", "WatchModelResources", "ProcessDyingModel", "RemoveModel"})
	c.Assert(destroyCloudCalls, jc.DeepEquals, []string{"ModelConfig", "CloudSpec", "Destroy"})
}

func (s *OldUndertakerSuite) sortCalls(c *gc.C, stub *testing.Stub) (mainCalls []string, destroyCloudCalls []string) {
	calls := stub.Calls()
	for _, call := range calls {
		switch call.FuncName {
		case "ModelConfig", "CloudSpec", "Destroy":
			destroyCloudCalls = append(destroyCloudCalls, call.FuncName)
		default:
			mainCalls = append(mainCalls, call.FuncName)
		}
	}
	return
}

func (s *OldUndertakerSuite) TestEnvironDestroyForceTimeoutZero(c *gc.C) {
	zero := time.Second * 0
	s.fix.info.Result.DestroyTimeout = &zero
	s.fix.info.Result.ForceDestroyed = true
	s.fix.dirty = true
	stub := s.fix.run(c, func(w worker.Worker) {
		err := workertest.CheckKilled(c, w)
		c.Assert(err, jc.ErrorIsNil)
	})
	stub.CheckCallNames(c, "WatchModel", "ModelInfo", "RemoveModel")
}

type UndertakerSuite struct{}

var _ = gc.Suite(&UndertakerSuite{})

func (s *UndertakerSuite) TestExitOnModelChanged(c *gc.C) {
	ctrl := gomock.NewController(c)
	defer ctrl.Finish()

	facade := NewMockFacade(ctrl)

	modelChanged := make(chan struct{}, 1)
	modelChanged <- struct{}{}
	modelResources := make(chan struct{}, 1)
	modelResources <- struct{}{}
	facade.EXPECT().WatchModel(gomock.Any()).DoAndReturn(func(context.Context) (watcher.NotifyWatcher, error) {
		return watchertest.NewMockNotifyWatcher(modelChanged), nil
	})
	facade.EXPECT().WatchModelResources(gomock.Any()).DoAndReturn(func(context.Context) (watcher.NotifyWatcher, error) {
		return watchertest.NewMockNotifyWatcher(modelResources), nil
	})
	facade.EXPECT().ModelConfig(gomock.Any()).Return(nil, nil)

	gomock.InOrder(
		facade.EXPECT().ModelInfo(gomock.Any()).Return(params.UndertakerModelInfoResult{
			Result: params.UndertakerModelInfo{
				Life:           life.Dying,
				ForceDestroyed: false,
			},
		}, nil),
		facade.EXPECT().ProcessDyingModel(gomock.Any()).Return(nil),
		facade.EXPECT().CloudSpec(gomock.Any()).DoAndReturn(func(_ context.Context) (cloudspec.CloudSpec, error) {
			modelChanged <- struct{}{}
			return cloudspec.CloudSpec{}, nil
		}),
		facade.EXPECT().ModelInfo(gomock.Any()).Return(params.UndertakerModelInfoResult{
			Result: params.UndertakerModelInfo{
				Life:           life.Dying,
				ForceDestroyed: true, // changed from false to true to cause worker to exit.
			},
		}, nil),
	)

	credentialAPI := NewMockCredentialAPI(ctrl)

	w, err := undertaker.NewUndertaker(undertaker.Config{
		Facade:        facade,
		CredentialAPI: credentialAPI,
		Logger:        loggertesting.WrapCheckLog(c),
		Clock:         testclock.NewDilatedWallClock(testing.ShortWait),
		NewCloudDestroyerFunc: func(ctx context.Context, op environs.OpenParams) (environs.CloudDestroyer, error) {
			return &waitDestroyer{}, nil
		},
	})
	c.Assert(err, jc.ErrorIsNil)

	// Ignore the error from the CheckKilled call, we'll check the specifics
	// later. We just want to make sure that the worker exits.
	_ = workertest.CheckKilled(c, w)

	err = w.Wait()
	c.Assert(err, gc.ErrorMatches, "model destroy parameters changed")
}

type waitDestroyer struct {
	environs.Environ
}

func (w *waitDestroyer) Destroy(ctx envcontext.ProviderCallContext) error {
	<-ctx.Done()
	return ctx.Err()
}
