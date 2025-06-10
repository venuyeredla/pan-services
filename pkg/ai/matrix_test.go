package ai

import (
	"fmt"
	"testing"
)

func TestMatrix(t *testing.T) {

	Matrix := [][]int{{2, 5, 8},
		{4, 0, -1}}

	/*Matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12}}

	/* Matrix2 := [][]int{{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9}}

	{13, 14, 15, 16}
	*/
	result := SpiralForm2(Matrix)
	for _, v := range result {
		fmt.Printf("%v, ", v)
	}

}

func TestMatrixRotation(t *testing.T) {

	Matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	RotateMatrixBy90(Matrix)
}
