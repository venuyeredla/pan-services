package array

// Range queries can be solved by Prefix sum array or Segment Trees and Binary Indexed trees.

func RangeQueriesSum(input []int, queries [][2]int) []int {
	sumArray := make([]int, len(input))
	output := make([]int, len(queries))
	preSum := 0
	for i, val := range input {
		sumArray[i] = preSum + val
	}
	for j := range queries {
		from := queries[j][0]
		to := queries[j][1]
		output[j] = sumArray[to] - sumArray[from-1]
	}
	return output
}

// Range queries  Building segment tree.
func RangeQueriesSegmentTree(input []int, queries [][2]int) []int {
	length := len(input)
	segmentTreeSize := length * (length - 1)
	segmentTree := make([]int, segmentTreeSize)
	output := make([]int, len(queries))

	buildSegmentTree(input, segmentTree, 0, 0, len(input))

	for j := range queries {
		from := queries[j][0]
		to := queries[j][1]
		output[j] = QuerySegemnt(segmentTree, from, to)
	}
	return output
}

func buildSegmentTree(input []int, segmentTree []int, sidx, left, right int) int {
	if left == right {
		segmentTree[left] = input[sidx]
		return input[left]
	} else {
		mid := getMid(left, right)
		leftSum := buildSegmentTree(input, segmentTree, sidx*2+1, left, mid)
		rightSum := buildSegmentTree(input, segmentTree, sidx*2+2, mid+1, right)
		segmentTree[sidx] = leftSum + rightSum
		return segmentTree[sidx]
	}
}

func QuerySegemnt(segement []int, from, to int) int {
	return 0

}

func getMid(l, r int) int {
	return l + (r-l)/2
}
