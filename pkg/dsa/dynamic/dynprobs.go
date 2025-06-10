package dynamic

import "github.com/venuyeredla/pan-services/pkg/dsa/utils"

func Rob(nums []int, idx int) int {
	if idx < 0 {
		return 0
	}
	includeSum := nums[idx] + Rob(nums, idx-2)
	excludeSum := Rob(nums, idx-1)

	return utils.MaxOf(includeSum, excludeSum)
}
