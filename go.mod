module flow_sim

go 1.13

require (
	github.com/RuiHirano/flow_sim/plugins v0.0.0-00010101000000-000000000000 // indirect
	github.com/RuiHirano/flow_sim/plugins/harmovis v0.0.0-00010101000000-000000000000 // indirect
	github.com/RuiHirano/flow_sim/plugins/higashiyama v0.0.0-00010101000000-000000000000 // indirect
	github.com/RuiHirano/flow_sim/plugins/rvo2 v0.0.0-00010101000000-000000000000 // indirect
	github.com/RuiHirano/flow_sim/plugins/sumo v0.0.0-00010101000000-000000000000 // indirect
	github.com/RuiHirano/flow_sim/util v0.0.0-00010101000000-000000000000 // indirect
	github.com/RuiHirano/rvo2-go/src/rvosimulator v0.0.0-20200707091306-e572a9b06cee // indirect
)

replace (
	github.com/RuiHirano/flow_sim/plugins => ./plugins
	github.com/RuiHirano/flow_sim/plugins/harmovis => ./plugins/harmovis
	github.com/RuiHirano/flow_sim/plugins/higashiyama => ./plugins/higashiyama
	github.com/RuiHirano/flow_sim/plugins/rvo2 => ./plugins/rvo2
	github.com/RuiHirano/flow_sim/plugins/sumo => ./plugins/sumo
	github.com/RuiHirano/flow_sim/util => ./util
)
