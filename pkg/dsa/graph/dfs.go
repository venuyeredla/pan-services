package graph

func validPath(n int, edges [][]int, source int, destination int) bool { // Adjacency list
	graph := make([][]int, n)
	for i := range graph {
		graph[i] = make([]int, 0)
	}
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	// BFS OR DFS
	visited := make([]bool, n)
	return DFS(graph, visited, source, destination)
}

func DFS(graph [][]int, visited []bool, source, destination int) bool {
	if source == destination {
		return true
	}
	if !visited[source] {
		visited[source] = true
		for _, newSource := range graph[source] {
			if DFS(graph, visited, newSource, destination) {
				return true
			}
		}
	}
	return false
}

/*
Sorrounded regions need to captured.

	    DFS? or BFS?
		0 < c(i,j) <R-1
		0 < c(i,j) <R-1

		Identfiying region
*/
func solve(board [][]byte) {
	ROWS := len(board)
	COLS := len(board[0])
	var dfsBoard func(i, j int)
	dfsBoard = func(i, j int) {
		if i < 0 || i == len(board) || j < 0 || j == len(board[0]) {
			return
		}
		if board[i][j] == 'X' || board[i][j] == 'R' {
			return
		}
		board[i][j] = 'R'
		dfsBoard(i, j+1)
		dfsBoard(i, j-1)
		dfsBoard(i+1, j)
		dfsBoard(i-1, j)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if (i == 0 || i == ROWS-1) || (j == 0 || j == COLS-1) {
				if board[i][j] != 'X' && !(board[i][j] == 'R') {
					dfsBoard(i, j)
				}
			}
		}
	}
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			if board[i][j] == 'R' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}

}
