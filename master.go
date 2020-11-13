package main
import (
	"log"
	"time"

	plugins "github.com/RuiHirano/flow_sim/plugins"
	harmovis "github.com/RuiHirano/flow_sim/plugins/harmovis"
	higashiyama "github.com/RuiHirano/flow_sim/plugins/higashiyama"
	rvo2 "github.com/RuiHirano/flow_sim/plugins/rvo2"
	sumo "github.com/RuiHirano/flow_sim/plugins/sumo"
	"github.com/RuiHirano/flow_sim/util"
)

var (
	realtime bool
)

func init(){
	realtime = true
}

type MasterProvider struct{
	Time uint64
	Agents []*util.Agent
	DataPlugins []plugins.DataPlugin
	SimPlugins []plugins.SimulationPlugin
	VisPlugins []plugins.VisualizationPlugin
}

func NewMasterProvider() *MasterProvider {
	mp := &MasterProvider{
		Time: 0,
		DataPlugins: []plugins.DataPlugin{},
		SimPlugins: []plugins.SimulationPlugin{},
		VisPlugins: []plugins.VisualizationPlugin{},
	}
	return mp
}

func (mp *MasterProvider) AddDataPlugins(plugins []plugins.DataPlugin) {
	mp.DataPlugins = append(mp.DataPlugins, plugins...)
}

func (mp *MasterProvider) AddSimPlugins(plugins []plugins.SimulationPlugin) {
	mp.SimPlugins = append(mp.SimPlugins, plugins...)
}

func (mp *MasterProvider) AddVisPlugins(plugins []plugins.VisualizationPlugin) {
	mp.VisPlugins = append(mp.VisPlugins, plugins...)
}

func (mp *MasterProvider) AddAgents(agents []*util.Agent) {
	mp.Agents = append(mp.Agents, agents...)
}

func (mp *MasterProvider) StartClock() {
	maxCycle := 100
	for i := 0; i < maxCycle; i++{
		t1 := time.Now()

		// realtime
		if realtime{
			// データプラグイン
			mp.handleDataPlugins()
			// シミュレーションプラグイン
			mp.handleSimPlugins()
			// 可視化プラグイン
			mp.handleVisPlugins()
			
		}else{
			// シミュレーションプラグイン
			mp.handleSimPlugins()
			
		}

		t2 := time.Now()
		duration := t2.Sub(t1).Milliseconds()
		interval := int64(1000) // 周期ms
		if realtime{
			if duration > interval {
				log.Printf("time cycle delayed... Duration: %d", duration)
			} else {
				// 待機
				log.Printf("Forward Clock! Time %d, Duration: %d ms, Wait: %d ms", mp.Time, duration, interval-duration)
				time.Sleep(time.Duration(interval-duration) * time.Millisecond)
			}
		}

		mp.Time += 1
	}
}


// データプラグインの処理
func (mp *MasterProvider) handleDataPlugins() {
	for _, plug := range mp.DataPlugins{
		plug.HandleAgents(mp.Agents)
	}
}

func (mp *MasterProvider) handleSimPlugins() {
	for _, plug := range mp.SimPlugins{
		plug.HandleAgents(mp.Agents)
	}
}

func (mp *MasterProvider) handleVisPlugins() {
	for _, plug := range mp.VisPlugins{
		plug.HandleAgents(mp.Agents)
	}
}


func main(){

	agents := util.GetMockAgents(100)
	// data plugins
	dataplugin := higashiyama.NewHigashiyamaPlugin()

	// sim plugins 
	rvo2plugin := rvo2.NewRVOPlugin()
	sumoplugin := sumo.NewRVOPlugin()

	// vis plugins
	visplugin := harmovis.NewHarmoVisPlugin()
	go visplugin.RunServer()

	master := NewMasterProvider()
	master.AddAgents(agents)
	master.AddDataPlugins([]plugins.DataPlugin{dataplugin})
	master.AddSimPlugins([]plugins.SimulationPlugin{rvo2plugin, sumoplugin})
	master.AddVisPlugins([]plugins.VisualizationPlugin{visplugin})
	master.StartClock()
}