// staticmodel
package compress

import (
	"fmt"
)

type SymDict struct {
	key   int
	count int
}

func Count(data []byte) []*SymDict {
	var counts [256]int
	for _, val := range data {
		counts[val] = counts[val] + 1
	}
	var symDict []*SymDict
	for i, val := range counts {
		if val != 0 {
			symDict = append(symDict, &SymDict{i, val})
		}
	}
	fmt.Println("Length=", len(symDict), " Capacity=", cap(symDict))
	return symDict

}
