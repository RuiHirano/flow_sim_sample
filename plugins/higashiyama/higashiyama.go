package plugins

import (
	"log"

	"github.com/RuiHirano/flow_sim/util"
)


type HigashiyamaPlugin struct {

}

func NewHigashiyamaPlugin() *HigashiyamaPlugin{
	return &HigashiyamaPlugin{}
}

func (rvo *HigashiyamaPlugin) HandleAgents(agents []*util.Agent) []*util.Agent{
	log.Printf("handle Agents higashiyama plugin")
	return agents
}
