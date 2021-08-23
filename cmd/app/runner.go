package main

import (
	"fmt"
	"github.com/lliryc/memego"
	"github.com/lliryc/memego/graph"
	"github.com/lliryc/memego/common"
)


func main() {
	const vNum = 5
	vertices, adjMap := memego.fullConnectedGraph(vNum)	
	generator := graph.GraphGenerator{Vertices: vertices, AdjMap: adjMap}
	policy := graph.TspPolicy {VerticesNum: vNum }
	firstGen := generator.Create(&policy)
	sim := 	common.Simulation {Generation: firstGen}
	res := sim.Run(5, &policy).(graph.Path)
	fmt.Printf("Best route is:")
	for _, v := range res.Vertices{
		fmt.Printf("%s", v.Id)
	}	
}