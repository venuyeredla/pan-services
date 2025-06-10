package backtracking

import (
	"fmt"
	"math"
	"testing"

	"github.com/venuyeredla/pan-services/pkg/dsa/utils"
)

func TestRatMaze(t *testing.T) {
	iniput := [][]int{
		{1, 0, 0, 0},
		{1, 1, 0, 1},
		{0, 1, 0, 0},
		{1, 1, 1, 1}}
	RatMaze(iniput)
}

func TestSubset(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	nofSubsets := int(math.Pow(2, float64(len(arr))))
	collector := utils.StringCollector(nofSubsets)
	PowerSet(arr, nofSubsets, collector)
	for _, v := range collector.Elements {
		fmt.Println(v)
	}
	fmt.Println("PowerSet generation by Backtracking")
	arr2 := []int{1, 2, 3} //, 3, 4
	collector2 := utils.StringCollector(8)
	PowerSetByBacktrack(arr2, collector2, 0, 2)
	for _, v := range collector2.Elements {
		fmt.Println(v)
	}
	//combinations(arr, 1, 2)
}

/*
	func TestPermuations(t *testing.T) {
		arr := []int{1, 2, 3, 4}
		size := maths.Factorial(len(arr))
		collector := utils.StringCollector(size)
		Permuations(arr, 0, len(arr), collector)
		for _, val := range collector.Elements {
			fmt.Println(val)
		}
	}
*/
func TestSubSequnces(t *testing.T) {
	subseq := SubSeqences("bbabcbcab") //bbabcbcab     abc
	for _, s := range subseq {
		if IsPalindrome(s) {
			fmt.Printf("%v - %v\n", s, len(s))
		}

	}
}

func IsPalindrome(str string /*, l, r int */) bool {
	/*if l >= r {
		return true
	} else if str[l:l+1] != str[r:r+1] {
		return false
	} else {
		return IsPalindrome(str, l+1, r-1)
	} */

	for l, r := 0, len(str)-1; l <= r; l, r = l+1, r-1 {
		if str[l] != str[r] {
			return false
		}
	}
	return true
}
