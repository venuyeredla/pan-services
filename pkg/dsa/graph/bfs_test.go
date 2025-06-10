package graph

import "testing"

func TestGraphBFS(t *testing.T) {
	edges := [][]int{{0, 1}, {1, 2}, {2, 0}}
	GraphWithEdges(3, false, edges)
}

func TestWordLadder(t *testing.T) {
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	restult := ladderLength("hit", "cog", wordList)
	if restult != 5 {
		t.Errorf("Failed Result = %v, Acutal = 5", restult)
	}
}
