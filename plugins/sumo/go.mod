module sumo

require (

)

replace (
	github.com/RuiHirano/flow_sim/plugins/sumo => ./plugins/sumo
	github.com/RuiHirano/flow_sim/util => ./util
)

go 1.13
