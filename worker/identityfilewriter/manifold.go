// Copyright 2016 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package identityfilewriter

import (
	"errors"

	"gopkg.in/juju/names.v2"
	"gopkg.in/juju/worker.v1"

	"github.com/juju/juju/agent"
	apiagent "github.com/juju/juju/api/agent"
	"github.com/juju/juju/api/base"
	"github.com/juju/juju/cmd/jujud/agent/engine"
	"github.com/juju/juju/state/multiwatcher"
	jworker "github.com/juju/juju/worker"
	"github.com/juju/juju/worker/dependency"
)

// ManifoldConfig defines the names of the manifolds on which a Manifold will depend.
type ManifoldConfig engine.AgentAPIManifoldConfig

// Manifold returns a dependency manifold that runs an identity file writer worker,
// using the resource names defined in the supplied config.
func Manifold(config ManifoldConfig) dependency.Manifold {
	typedConfig := engine.AgentAPIManifoldConfig(config)
	return engine.AgentAPIManifold(typedConfig, newWorker)
}

// newWorker trivially wraps NewWorker for use in a engine.AgentAPIManifold.
func newWorker(a agent.Agent, apiCaller base.APICaller) (worker.Worker, error) {
	cfg := a.CurrentConfig()

	// Grab the tag and ensure that it's for a machine.
	tag, ok := cfg.Tag().(names.MachineTag)
	if !ok {
		return nil, errors.New("this manifold may only be used inside a machine agent")
	}

	// Get the machine agent's jobs.
	entity, err := apiagent.NewState(apiCaller).Entity(tag)
	if err != nil {
		return nil, err
	}

	var isModelManager bool
	for _, job := range entity.Jobs() {
		if job == multiwatcher.JobManageModel {
			isModelManager = true
			break
		}
	}

	if !isModelManager {
		return nil, dependency.ErrMissing
	}

	return NewWorker(cfg)
}

var NewWorker = func(agentConfig agent.Config) (worker.Worker, error) {
	inner := func(<-chan struct{}) error {
		return agent.WriteSystemIdentityFile(agentConfig)
	}
	return jworker.NewSimpleWorker(inner), nil
}
