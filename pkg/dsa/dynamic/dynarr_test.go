package dynamic

import (
	"fmt"
	"math"
	"testing"

	"github.com/venuyeredla/pan-services/pkg/dsa/utils"
)

func TestSubsetSum(t *testing.T) {
	input := []int{3, 2, 7, 1}
	result := SubsetSum(input, 6)
	resp := SubsetSumD(input, 6)
	if result != resp {
		t.Error("Failed to pass test")
		t.FailNow()
	}
}

func TestLargest(t *testing.T) {
	arr := []int{2, -3, 4, -1, -2, 1, 5, -3}
	LargestSumSubArray(arr)
}

func TestLargeMonotonic(t *testing.T) {
	arr := []int{3, 10, 2} // 3, 10, 2, 1, 20
	result := longIncreasingSubseq(arr, 0, 0)
	if result != 3 {
		t.Error("Failed to pass test")
		t.FailNow()
	}
}

// First increase J to increase sum beyond or equal to sum.
// Second increase i to minimize the length
func minSubArrayLen(target int, nums []int) int {
	// nums[i:j]>=traget.  return j-i
	ms, sum := math.MaxInt, 0
	for i, j := 0, 0; j < len(nums); j++ {
		sum = sum + nums[j]
		if sum >= target {
			ms = utils.Min(ms, j-i+1)
			for sum-nums[i] >= target {
				sum = sum - nums[i]
				i = i + 1
				ms = utils.Min(ms, j-i+1)
			}
		}
	}
	if ms == math.MaxInt {
		return 0
	} else {
		return ms
	}
}

func TestFuck(t *testing.T) {
	nums := []int{10, 2, 3}
	min := minSubArrayLen(11, nums)
	fmt.Printf("value =%v \n", min)
}
