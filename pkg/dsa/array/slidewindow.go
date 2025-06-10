package array

import (
	"math"

	"github.com/venuyeredla/pan-services/pkg/dsa/utils"
)

func shortestSubarray(nums []int, k int) int {
	// nums[i:j]>=traget.  return j-i
	ms, sum := math.MaxInt, 0
	for i, j := 0, 0; j < len(nums); j++ {
		sum = sum + nums[j]
		if sum >= k {
			ms = utils.Min(ms, j-i+1)
			for sum-nums[i] >= k {
				sum = sum - nums[i]
				i = i + 1
				ms = utils.Min(ms, j-i+1)
			}
		}
	}
	if ms == math.MaxInt {
		return -1
	} else {
		return ms
	}

}
