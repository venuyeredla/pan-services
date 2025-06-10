package stack_queue

import (
	"fmt"
	"testing"

	"github.com/venuyeredla/pan-services/pkg/dsa/utils"
)

func heap1(min bool) *PriorityQueue {
	h := NewPQ(12, min)
	if min {
		h.list = []PqEntry{
			{"x", 3},
			{"t", 7},
			{"o", 5},
			{"g", 18},
			{"s", 16},
			{"m", 19},
			{"n", 14},
			{"b", 28},
			{"e", 27},
			{"r", 29},
			{"a", 26},
			{"i", 23},
		}
	} else {
		h.list = []PqEntry{
			{"x", 25},
			{"t", 22},
			{"o", 20},
			{"g", 18},
			{"s", 16},
			{"m", 19},
			{"n", 14},
			{"b", 8},
			{"e", 7},
			{"r", 9},
			{"a", 6},
			{"i", 3},
		}
	}
	return h
}

func TestMaxFixUp(x *testing.T) {
	t := (*utils.T)(x)
	h := heap1(false)
	t.Log(h)
	h.list[10].Priority = 30
	h.fixUp(10)
	t.Log(h)
	t.Assert(h.list[0].Item.(string) == "a", "heap did not start with {a 30} %v", h)
}

func TestMinFixUp(x *testing.T) {
	t := (*utils.T)(x)
	h := heap1(true)
	t.Log(h)
	h.list[10].Priority = 1
	h.fixUp(10)
	t.Log(h)
	t.Assert(h.list[0].Item.(string) == "a", "heap did not start with {a 1} %v", h)
}

func TestMaxFixDown(x *testing.T) {
	t := (*utils.T)(x)
	h := heap1(false)
	t.Log(h)
	h.list[0].Priority = 0
	h.fixDown(0)
	t.Log(h)
	t.Assert(h.list[0].Item.(string) == "t", "heap did not start with {t 22} %v", h)
	t.Assert(h.list[7].Item.(string) == "x", "heap[8] != {x 0} %v", h)
}

func TestMinFixDown(x *testing.T) {
	t := (*utils.T)(x)
	h := heap1(true)
	t.Log(h)
	h.list[0].Item = 30
	h.fixDown(0)
	t.Log(h)
	t.Assert(h.list[0].Item.(string) == "o", "heap did not start with {o 5} %v", h)
	t.Assert(h.list[6].Item.(string) == "x", "heap[8] != {n 30} %v", h)
}

func TestPushMax(x *testing.T) {
	t := (*utils.T)(x)
	h := NewPQ(12, false)
	input := map[string]int{
		"g": 18,
		"e": 3,
		"i": 6,
		"a": 25,
		"x": 22,
		"t": 14,
		"n": 8,
		"m": 19,
		"r": 20,
	}
	for key, value := range input {
		h.Push(PqEntry{key, value})
	}

	t.Log(h)
	t.AssertNil(h.Verify())
}

/* PQ testing */

func TestSorting(t *testing.T) {
	//t.Skip()
	//input := GenArray(5, 20)
	input := []int{10, 80, 30, 90, 40, 50, 70}
	//expected := []int{10, 30, 40, 50, 70, 80, 90}
	HeapSort(input)
}

func TestKmost(t *testing.T) {
	arr := []int{7, 10, 11, 5, 2, 5, 5, 7, 11, 8, 9}
	result := KMostOccurance(arr, 4)
	fmt.Println("Result = %%V", result)
}

func TestKthSumSubArr(t *testing.T) {
	arr := []int{10, -10, 20, -40} // 20. 15, 14
	result := KthLargestSumSubArray(arr, 6)
	fmt.Println("Result = %V", result)
}

/*
func TestRearrange(t *testing.T) {
	input := "aaabc"
	result := rearangeString(input)
	fmt.Println("Result = %V", result)
}
*/
