package graph

import "fmt"

// Number of vertices in the graph
const VP int = 5

// A utility function to find the vertex with minimum
// key value, from the set of vertices not yet included
// in MST
func minKey(key []int, mstSet []bool) int {
	// Initialize min value
	min := 4294967295 // Integer.MAX_VALUE
	min_index := -1

	for v := 0; v < VP; v++ {
		if mstSet[v] == false && key[v] < min {
			min = key[v]
			min_index = v
		}
	}

	return min_index
}

// A utility function to print the constructed MST
// stored in parent[]
func printMST(parent []int, graph [][]int) {
	fmt.Println("Edge \tWeight")
	for i := 1; i < VP; i++ {
		fmt.Printf("%v  _ %v \t %v", parent[i], i, graph[i][parent[i]])
	}

}

// Function to construct and print MST for a graph
// represented using adjacency matrix representation
func primMST(graph [][]int) {
	// Array to store constructed MST
	parent := make([]int, VP)

	// Key values used to pick minimum weight edge in
	// cut
	key := make([]int, VP)

	// To represent set of vertices included in MST
	mstSet := make([]bool, VP)

	// Initialize all keys as INFINITE
	for i := 0; i < VP; i++ {
		key[i] = 4294967295 // Integer.MAX_VALUE
		mstSet[i] = false
	}

	// Always include first 1st vertex in MST.
	key[0] = 0 // Make key 0 so that this vertex is
	// picked as first vertex
	parent[0] = -1 // First node is always root of MST

	// The MST will have V vertices
	for count := 0; count < V-1; count++ {
		// Pick thd minimum key vertex from the set of
		// vertices not yet included in MST
		var u int = minKey(key, mstSet)

		// Add the picked vertex to the MST Set
		mstSet[u] = true

		// Update key value and parent index of the
		// adjacent vertices of the picked vertex.
		// Consider only those vertices which are not
		// yet included in MST
		for v := 0; v < VP; v++ {
			if graph[u][v] != 0 && mstSet[v] == false && graph[u][v] < key[v] {
				parent[v] = u
				key[v] = graph[u][v]
			}
		}

		// graph[u][v] is non zero only for adjacent
		// vertices of m mstSet[v] is false for
		// vertices not yet included in MST Update
		// the key only if graph[u][v] is smaller
		// than key[v]
	}

	// print the constructed MST
	printMST(parent, graph)
}
