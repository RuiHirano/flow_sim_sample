module harmovis

replace (
	github.com/RuiHirano/flow_sim/plugins/harmovis => ./plugins/harmovis
	github.com/RuiHirano/flow_sim/util => ./util
)

go 1.13

require github.com/mtfelian/golang-socketio v1.5.2 // indirect
