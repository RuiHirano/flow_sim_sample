module higashiyama

require (

)

replace (
	github.com/RuiHirano/flow_sim/plugins/higashiyama => ./plugins/higashiyama
	github.com/RuiHirano/flow_sim/util => ./util
)

go 1.13
