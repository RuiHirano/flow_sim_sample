package plugins

import (
	"log"

	"github.com/RuiHirano/flow_sim/util"
)

type RVOPlugin struct {

}

func NewRVOPlugin() *RVOPlugin{
	return &RVOPlugin{}
}

func (rvo *RVOPlugin) HandleAgents(agents []*util.Agent) []*util.Agent{
	// 他のシミュレータデータ、障害物の処理
	// start simulation
	log.Printf("handle Agents rvo plugin")
	nextAgents := rvo.startSimulation(agents)
	return nextAgents
}

func (rvo *RVOPlugin) startSimulation(agents []*util.Agent) []*util.Agent{

	/*timeStep := 0.1
	neighborDist := 0.00003 // どのくらいの距離の相手をNeighborと認識するか?Neighborとの距離をどのくらいに保つか？ぶつかったと認識する距離？
	maxneighbors := 3       // 周り何体を計算対象とするか
	timeHorizon := 1.0
	timeHorizonObst := 1.0
	radius := 0.00001  // エージェントの半径
	maxSpeed := 0.0004 // エージェントの最大スピード
	sim := rvo.NewRVOSimulator(timeStep, neighborDist, maxneighbors, timeHorizon, timeHorizonObst, radius, maxSpeed, &rvo.Vector2{X: 0, Y: 0})*/

	// scenario設定
	//rvo2route.SetupScenario()

	// Stepを進める
	//sim.DoStep()

	// 管理エリアのエージェントのみを抽出
	/*nextAgents := make([]*util.Agent, 0)
	for _, agentInfo := range agents {

		nextAgent := &util.Agent{
			agentInfo.ID,
			agentInfo.Name,
		}

		nextAgents = append(nextAgents, nextAgent)
	}*/
	nextAgents := agents

	return nextAgents
}