package array

func productExceptSelf(nums []int) []int {
	length := len(nums)
	if length < 1 {
		return []int{}
	}
	suffixProduct := make([]int, len(nums))
	suffixProduct[length-1] = nums[length-1]
	for k := length - 2; k >= 0; k-- {
		suffixProduct[k] = suffixProduct[k+1] * nums[k]
	}

	answer := make([]int, len(nums))
	answer[0] = suffixProduct[1]
	leftProduct := nums[0]
	for i := 1; i < length-1; i++ {
		answer[i] = leftProduct * suffixProduct[i+1]
		leftProduct = leftProduct * nums[i]
	}
	answer[length-1] = leftProduct
	return answer
}

// Using cyclic sort for sorting numbers in given range.
func findDuplicate(nums []int) int {
	occrance := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _, ok := occrance[nums[i]]; ok {
			return nums[i]
		}
		occrance[nums[i]] = true
	}
	return -1
}

/*
Algoirthm:

 1. Note down how far you can move from current step
 2. For each positon take maxumm jumps you make
 3. When you are exceeding previous jump range increase count
*/
func minJumpsRequred(nums []int) int {
	var max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	minJumps := 0
	currentJumpRange := 0
	maxJump := 0
	for i := 0; i < len(nums); i++ {
		if maxJump <= i && nums[i] <= 0 {
			return -1
		}
		maxJump = max(maxJump, i+nums[i])
		if currentJumpRange == i {
			minJumps++
			currentJumpRange = i + nums[i]
		}
	}
	return minJumps
}
