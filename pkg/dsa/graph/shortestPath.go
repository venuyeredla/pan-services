package graph

import "fmt"

// A utility function to find the vertex with minimum
// distance value, from the set of vertices not yet
// included in shortest path tree
const V int = 9

func minDistance(dist []int, sptSet []bool) int {
	// Initialize min value
	var min int = 4294967295 // Integer.MAX_VALUE
	var min_index int = -1

	for v := 0; v < V; v++ {
		if sptSet[v] == false && dist[v] <= min {
			min = dist[v]
			min_index = v
		}
	}

	return min_index
}

// A utility function to print the constructed distance
// array
func printSolution(dist []int) {
	fmt.Println("Vertex \t\t Distance from Source")
	for i := 0; i < V; i++ {
		fmt.Printf("%d \t\t %d", i, dist[i])
	}

}

// Function that implements Dijkstra's single source
// shortest path algorithm for a graph represented using
// adjacency matrix representation
func dijkstra(graph [][]int, src int) {

	dist := make([]int, V)
	// The output array.
	// dist[i] will hold
	// the shortest distance from src to i

	// sptSet[i] will true if vertex i is included in
	// shortest path tree or shortest distance from src
	// to i is finalized
	sptSet := make([]bool, V)

	// Initialize all distances as INFINITE and stpSet[]
	// as false
	for i := 0; i < V; i++ {
		dist[i] = 4294967295 // Integer.MAX_VALUE
		sptSet[i] = false
	}

	// Distance of source vertex from itself is always 0
	dist[src] = 0

	// Find shortest path for all vertices
	for count := 0; count < V-1; count++ {

		// Pick the minimum distance vertex from the set
		// of vertices not yet processed. u is always
		// equal to src in first iteration.
		var u int = minDistance(dist, sptSet)

		// Mark the picked vertex as processed
		sptSet[u] = true

		// Update dist value of the adjacent vertices of
		// the picked vertex.
		for v := 0; v < V; v++ {
			if !sptSet[v] && graph[u][v] != 0 && dist[u] != 4294967295 && dist[u]+graph[u][v] < dist[v] {
				dist[v] = dist[u] + graph[u][v]
			}

		}

		// Update dist[v] only if is not in sptSet,
		// there is an edge from u to v, and total
		// weight of path from src to v through u is
		// smaller than current value of dist[v]

	}

	// print the constructed distance array
	printSolution(dist)
}
