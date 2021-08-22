package main

import (
	. "github.com/lliryc/memego/graph"
	"fmt"
	"math/rand"
)

func fullConnectedGraph(n int) ([] Vertex, map[string]map[string]float32){
	vertices := make([]Vertex, n)
	for i,_ := range vertices {
		vertices[i] = Vertex{Id: fmt.Sprintf("%d",i)}
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
	generator := GraphGenerator{Vertices: vertices, AdjMap: adjMap}
	policy := TspPolicy {VerticesNum: vNum }
	firstGen := generator.Create(&policy)
	sim := 	memego.Simulation {Generation: firstGen}
	res := sim.Run(5, &policy).(Path)
	fmt.Printf("Best route is:")
	for _, v := range res.Vertices{
		fmt.Printf("%s", v.Id)
	}	
}