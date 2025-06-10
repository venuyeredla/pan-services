package backtracking

import (
	"fmt"
	"strings"

	"github.com/venuyeredla/pan-services/pkg/dsa/utils"
)

func PowerSet(arr []int, nofSubsets int, collector *utils.Collector) {
	for idx := 0; idx < nofSubsets; idx++ {
		i := -1
		var sb strings.Builder
		sb.WriteString("{")
		num := idx
		for num > 0 {
			i += 1
			if num&1 == 1 {
				fmt.Fprintf(&sb, "%d", arr[i])
			}
			num = num >> 1
		}
		sb.WriteString("}")
		collector.Append(sb.String())
	}
}

// Sets containing element and not containg element.
// // 1,2,3 -> 1,2,3
//     and  2,3

func PowerSetByBacktrack(arr []int, collector *utils.Collector, l, r int) {
	if len(arr) == 1 {
		collector.Append(utils.Printable(arr, l, r))
	} else if len(arr) > 1 {
		collector.Append(utils.Printable(arr, l, r))
		subsize := r - l
		for i := l + 1; i < len(arr); i++ {
			include := make([]int, 0, subsize)
			for j := l; j < len(arr); j++ {
				if i != j {
					include = append(include, arr[j])
				}
			}
			//	fmt.Printf("include =%v", include)
			PowerSetByBacktrack(include, collector, 0, len(include)-1)
		}

		exclude := make([]int, subsize)
		copy(exclude, arr[l+1:])
		if len(exclude) > 1 {
			//fmt.Printf("Exclude =%v", exclude)
			PowerSetByBacktrack(exclude, collector, 0, len(exclude)-1)
		}
	}
}

// Assumption arr has unique elements {1,2,3}
func combinations(arr []int, to, size int) {
	/* for i := 0; i < len(arr); i++ {
		Swap(arr, to, i)
		PrintArr(arr, 0, to)
		Swap(arr, i, to)
	} */
	selection := make([]int, size) //[2]int{-1, -1}
	for i := 0; i < len(arr); i++ {
		selection[0] = arr[i]
		for j := i + 1; j < len(arr); j++ {
			selection[1] = arr[j]
			utils.Printable(selection[:], 0, 1)
		}
	}
}
