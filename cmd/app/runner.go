package main

import (
	"fmt"
	"math/rand"
	"github.com/lliryc/memego/graph"
	"github.com/lliryc/memego/common"
)

func fullConnectedGraph(n int) ([] graph.Vertex, map[string]map[string]float32){
	vertices := make([]graph.Vertex, n)
	for i,_ := range vertices {
		vertices[i] = graph.Vertex{Id: fmt.Sprintf("%d",i)}
	}
	adjMap := make(map[string]map[string]float32) 
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			dist := rand.Float32()
			adjMap[fmt.Sprintf("%d",i)][fmt.Sprintf("%d",j)] = dist
			adjMap[fmt.Sprintf("%d",j)][fmt.Sprintf("%d",i)] = dist
		}
	}
	return vertices, adjMap 
}

func main() {
	const vNum = 5
	vertices, adjMap := fullConnectedGraph(vNum)	
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