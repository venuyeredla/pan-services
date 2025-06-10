package utils

import (
	"fmt"
	"math"
	mathrand "math/rand"
	"strings"
	"time"
)

func Printable(arr []int, l, r int) string {
	if l <= r {
		var sb strings.Builder
		sb.WriteString("{")
		for k := l; k <= r; k++ {
			fmt.Fprintf(&sb, "%d", arr[k])
			if k < r {
				sb.WriteString(",")
			}
		}
		sb.WriteString("}")
		s := sb.String()
		fmt.Println(s)
		return s
	}
	return ""
}

func GenArray(size int, max int) []int {
	// balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}  // static arry with intialization of values.
	//var generated [5]int // fixed array size declaration
	var generated []int //If size don't mentined it will become slice. Before using slice need to intilaize.
	mathrand.Seed(time.Now().UnixMilli())
	generated = make([]int, size)
	for i := 0; i < size; i++ {
		generated[i] = mathrand.Intn(20)
	}
	return generated
}

func GetMatrix(rows, cols int) [][]int {
	sm := make([][]int, rows)
	for i := 0; i < rows; i++ {
		sm[i] = make([]int, cols)
	}
	return sm
}

func Minimum(i, j, k int) int {
	return int(math.Min(float64(math.Min(float64(i), float64(j))), float64(k)))
}

func MaxOf(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
