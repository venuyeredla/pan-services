package compress

import (
	"fmt"
	"math"
)

func AbsCompress(bitStr string) {
	fmt.Println("Building model by calculating frequencies.")
	buildABSModel(bitStr)
}

func buildABSModel(bitStr string) {
	ones := 0
	zeros := 0
	for _, bit := range bitStr {
		if bit == 48 {
			zeros++
		} else if bit == 49 {
			ones++
		}
	}

	total := ones + zeros
	fmt.Println("Total:", total, " 1's:", ones, "	0's:", zeros)
	fmt.Println()

	p1 := float64(ones) / float64(total)
	p0 := 1 - p1
	fmt.Println(" p0=", p0, " , p1=", p1)
	var stable [35][2]int //= new int[24][2];
	var newState float64 = 0

	for i := 1; i <= 14; i++ {
		for j := 0; j < 2; j++ {
			cstate := float64(i)
			if j == 0 {
				newState = math.Floor(cstate / p0)
				stable[int(newState)][1] = 0
			} else if j == 1 {
				newState = math.Floor(cstate / p1)
				stable[int(newState)][1] = 1
			}
			stable[int(newState)][0] = i
		}
	}
	fmt.Println("State table:")
	for i := 1; i < 25; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	for j := 1; j < 25; j++ {
		fmt.Print(stable[j][0], " ")
		if j > 9 && stable[j][0] < 10 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
	for k := 1; k < 25; k++ {
		fmt.Print(stable[k][1], " ")
		if k > 9 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
