// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package trace

import (
	"context"

	"github.com/DataDog/datadog-agent/pkg/trace/agent"
	"github.com/DataDog/datadog-agent/pkg/trace/config"
	"github.com/DataDog/datadog-agent/pkg/util/log"
)

// ServerlessTraceAgent represents a trace agent in a serverless context
type ServerlessTraceAgent struct {
	ta     *agent.Agent
	cancel context.CancelFunc
}

type LoadConfig struct {
	Path string
}

type Load interface {
	Load() (*config.AgentConfig, error)
}

func (l *LoadConfig) Load() (*config.AgentConfig, error) {
	return config.Load(l.Path)
}

// Start starts the agent
func (s *ServerlessTraceAgent) Start(enabled bool, loadConfig Load, waitingChan chan bool) {
	if enabled {
		tc, confErr := loadConfig.Load()
		if confErr != nil {
			log.Errorf("Unable to load trace agent config: %s", confErr)
		} else {
			context, cancel := context.WithCancel(context.Background())
			tc.Hostname = ""
			tc.SynchronousFlushing = true
			s.ta = agent.NewAgent(context, tc)
			s.cancel = cancel
			go func() {
				s.ta.Run()
			}()
		}
	}
	waitingChan <- true
}

// Get returns the trace agent instance
func (s *ServerlessTraceAgent) Get() *agent.Agent {
	return s.ta
}

// Stop stops the trace agent
func (s *ServerlessTraceAgent) Stop() {
	s.cancel()
}
