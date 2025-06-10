package ai

import "fmt"

func Graph(inputs [][]int) {

	matrix := NewMatrix(10, 10, 0)
	for _, value := range inputs {
		matrix.Set(10-value[0], 10-value[1], 1)
	}

	fmt.Print("Graph : \n")
	for i, row := range matrix.Data {
		fmt.Print("|")
		for _, col := range row {
			if col != 0 {
				fmt.Print("* ")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Printf("\n|%v\n", 10-i)
	}
	for i := 0; i < matrix.Shape().Cols; i++ {
		fmt.Printf("%v---", i)
	}
	fmt.Print("\n")
}
