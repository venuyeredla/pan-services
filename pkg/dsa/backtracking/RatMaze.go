package backtracking

import "fmt"

var mazX, mazY int

func RatMaze(input [][]int) {
	mazY = len(input) - 1
	mazX = len(input[0]) - 1
	output := make([][]int, 4, 4)
	for i := 0; i < len(output); i++ {
		arr := make([]int, 4, 4)
		for j := 0; j < len(arr); j++ {
			arr[j] = 0
		}
		output[i] = arr
	}
	moveRat(input, output, 0, 0)
	fmt.Printf("Out put is %v", output)
}

func moveRat(input [][]int, output [][]int, y, x int) bool {

	if x == mazX && y == mazY {
		output[y][x] = 1
		return true
	}
	if input[y][x] != 1 {
		return false
	}

	if output[y][x] == -1 {
		return false
	}
	output[y][x] = 1
	if (x + 1) <= mazX {
		xres := moveRat(input, output, y, x+1)
		if !xres {
			output[y][x+1] = -1
		}
	}

	if (y + 1) <= mazY {
		yresult := moveRat(input, output, y+1, x)
		if !yresult {
			output[y+1][x] = -1
		}
	}

	return true
}
