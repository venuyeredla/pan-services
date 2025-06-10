package backtracking

import (
	"fmt"
	"strings"

	"github.com/venuyeredla/pan-services/pkg/dsa/utils"
)

func Permuations(arr []int, left, right int, collector *utils.Collector) {
	if left == right {
		var sb strings.Builder
		for k := 0; k <= right-1; k++ {
			fmt.Fprintf(&sb, "%d", arr[k])
		}
		collector.Append(sb.String())
	} else {
		for i := left; i < right; i++ {
			arr[i], arr[left] = arr[left], arr[i]
			Permuations(arr, left+1, right, collector)
			arr[i], arr[left] = arr[left], arr[i]
		}
	}
}

/*
Input : abc
Output : a, ab, abc, ac   b, bc, c

	inclusion and eclusion principle. Backtracking and recurison
*/
func SubSeqences(str string) []string {
	collector := utils.StringCollector(8)
	subSeq(str, "", 0, len(str), collector)
	return collector.Elements
}

func subSeq(str, sub string, left, right int, collector *utils.Collector) {
	if str == "" || left == right {
		return
	}
	newSub := sub + string(str[left])
	collector.Append(newSub)
	subSeq(str, newSub, left+1, right, collector)
	subSeq(str, sub, left+1, right, collector)
}
