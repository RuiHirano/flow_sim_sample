package plugins

import (
	"github.com/RuiHirano/flow_sim/util"
)


type DataPlugin interface {
	HandleAgents(agents []*util.Agent) []*util.Agent
}

type SimulationPlugin interface {
	HandleAgents(agents []*util.Agent) []*util.Agent
}

type VisualizationPlugin interface {
	HandleAgents(agents []*util.Agent)
	RunServer()
}