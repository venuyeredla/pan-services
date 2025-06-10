package graph

import (
	"fmt"
	"testing"
)

func TestPathExist(t *testing.T) {
	edges := [][]int{{4, 3}, {1, 4}, {4, 8}, {1, 7}, {6, 4}, {4, 2}, {7, 4}, {4, 0}, {0, 9}, {5, 4}}
	out := validPath(10, edges, 5, 9)
	fmt.Println(out)

}

func TestSolveBoard(t *testing.T) {
	//board := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	board := [][]byte{{'O', 'O'},
		{'O', 'O'}}
	solve(board)
}
