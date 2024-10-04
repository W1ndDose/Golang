package main

import (
	"fmt"
)

type Graph struct {
	vertices map[int][]int
}

func (g *Graph) addEdge(v, w int) {
	if g.vertices == nil {
		g.vertices = make(map[int][]int)
	}
	g.vertices[v] = append(g.vertices[v], w)
	g.vertices[w] = append(g.vertices[w], v)
}

func (g *Graph) bfs(start int) {
	visited := make(map[int]bool)
	queue := []int{start}

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]

		if !visited[vertex] {
			fmt.Printf("%d ", vertex)
			visited[vertex] = true

			for _, neighbor := range g.vertices[vertex] {
				if !visited[neighbor] {
					queue = append(queue, neighbor)
				}
			}
		}
	}
}

func (g *Graph) dfs(start int, visited map[int]bool) {
	visited[start] = true
	fmt.Printf("%d ", start)

	for _, neighbor := range g.vertices[start] {
		if !visited[neighbor] {
			g.dfs(neighbor, visited)
		}
	}
}

func main() {
	g := Graph{}

	g.addEdge(1, 2)
	g.addEdge(1, 3)
	g.addEdge(2, 3)
	g.addEdge(3, 4)
	g.addEdge(4, 5)
	g.addEdge(5, 6)

	fmt.Println("Breadth-First Traversal starting from vertex 1:")
	g.bfs(1)
	fmt.Println()

	fmt.Println("Depth-First Traversal starting from vertex 1:")
	visited := make(map[int]bool)
	g.dfs(1, visited)
	fmt.Println()
}
