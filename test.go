package main

import "log"

type DataPlugin interface {
	HandleAgents()
}

type RVOPlugin struct {

}

func NewRVOPlugin() *RVOPlugin{
	return &RVOPlugin{}
}

func (rvo *RVOPlugin) HandleAgents(){
	log.Printf("handle agents")
	rvo.HandleAgents2()
}

func (rvo *RVOPlugin) HandleAgents2(){
	log.Printf("handle agents2")
}

func main(){
	var dataplugin DataPlugin 
	dataplugin = NewRVOPlugin()
	dataplugin.HandleAgents()
}