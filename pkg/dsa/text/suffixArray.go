package text

import (
	"fmt"
	"sort"
)

type SuffixIdx struct {
	suffix string
	index  int
}

var suffixArray []*SuffixIdx

func PatMatchSuffixArray(input, pattern string) int {
	suffixArray = make([]*SuffixIdx, 0, len(input))
	for i := 0; i < len(input); i++ {
		suffixArray = append(suffixArray, &SuffixIdx{suffix: input[i:], index: i})
	}
	sort.Slice(suffixArray, func(i, j int) bool {
		return suffixArray[i].suffix < suffixArray[j].suffix
	})

	idx := sort.Search(len(suffixArray), func(i int) bool {
		return suffixArray[i].suffix >= pattern
	})
	for _, sufIdx := range suffixArray {
		fmt.Printf("%v - %v \n", sufIdx.suffix, sufIdx.index)
	}

	if idx == -1 {
		return -1
	} else {
		return suffixArray[idx].index
	}
}
