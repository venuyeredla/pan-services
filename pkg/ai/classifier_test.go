package ai

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	t.Skip()
	inputs := [][]int{
		{1, 3},
		{2, 5},
		{3, 5},
		{2, 1},
		{4, 2},
		{5, 2},
	}

	Graph(inputs)
}

func TestLinearClassifier(t *testing.T) {

	inputs := [][]int{
		{1, 3},
		{2, 5},
		{3, 5},
		{2, 1},
		{4, 2},
		{5, 2},
	}
	targets := []int{-1, -1, -1, 1, 1, 1}
	lm := LinearModel{}
	lm.train(inputs, targets)
	for i, value := range inputs {
		predicted := lm.pridict(value)
		fmt.Printf("Expected = % v and predicted = %v \n", targets[i], predicted)
	}

	predicted := lm.pridict([]int{5, 2})
	fmt.Printf(" New value Expected = % v and predicted = %v \n", -1, predicted)
}
