package memego

import (
	"github.com/lliryc/memego/graph"
	"fmt"
	"math/rand"
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