module plugins

require (

)

replace (
	github.com/RuiHirano/flow_sim/plugins => ./plugins
	github.com/RuiHirano/flow_sim/util => ./util
)

go 1.13
