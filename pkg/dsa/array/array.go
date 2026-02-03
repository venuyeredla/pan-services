package array

import (
	"fmt"

	"github.com/venuyeredla/pan-services/pkg/dsa/utils"
)

type Salgo byte

const (
	Bubble Salgo = iota
	Selection
	Insertion
	Heap
	Quick
	Merge
	Counting
	Radix
	Bucket
)

// arrange numbers in ai<aj  where i<j
func Sort(input []int, algo Salgo) {
	length := len(input)
	switch algo {
	case Bubble:
		for i := range length {
			for j := 0; j < length-i-1; j++ {
				if input[j] > input[j+1] {
					input[j], input[j+1] = input[j+1], input[j]
				}
			}
		}

	case Selection:
		for i := range length {
			minIdx := i
			for j := i + 1; j < length; j++ {
				if input[minIdx] > input[j] {
					minIdx = j
				}
			}
			if minIdx != i {
				input[minIdx], input[i] = input[i], input[minIdx]
			}
		}

	case Insertion:
		for i := 1; i < len(input); i++ {
			j := i - 1
			temp := input[i]
			for j >= 0 && temp < input[j] {
				input[j+1] = input[j]
				j--
			}
			j++
			if j != i {
				input[j] = temp
			}

		}

	case Quick:
		QuickSort(input, 0, len(input)-1)

	case Merge:
		MergSort(input, 0, length-1)
	}
}

func QuickSort(input []int, left, right int) {
	if left < right {
		pivot := getPivot(input, left, right)
		fmt.Printf("QuickSort Left=[%v, %v], Right=[%v,%v] \n", left, pivot-1, pivot+1, right)
		QuickSort(input, left, pivot-1)
		QuickSort(input, pivot+1, right)
	}

}

func getPivot(input []int, left, right int) int {
	pivot := input[right]
	i := left - 1
	for j := left; j < right; j++ {
		if input[j] < pivot {
			i = i + 1
			input[i], input[j] = input[j], input[i]
		}
	}
	input[i+1], input[right] = input[right], input[i+1]
	return (i + 1)
}

func MergSort(arr []int, l, r int) {
	if l < r {
		// Find the middle point
		m := l + (r-l)/2
		// Sort first and second halves
		MergSort(arr, l, m)
		MergSort(arr, m+1, r)
		// Merge the sorted halves
		merge(arr, l, m, r)
	}
}

func merge(arr []int, l, m, r int) {
	//var L []int = make([]int, n1)
	LT := arr[l : m+1]
	RT := arr[m+1 : r+1]
	L := make([]int, len(LT))
	R := make([]int, len(RT))
	copy(L, LT[:])
	copy(R, RT[:])
	i, j := 0, 0
	k := l
	for i < len(L) && j < len(R) {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}

	/* Copy remaining elements of L[] if any */
	for i < len(L) {
		arr[k] = L[i]
		i++
		k++
	}

	/* Copy remaining elements of R[] if any */
	for j < len(R) {
		arr[k] = R[j]
		j++
		k++
	}
}

// Since array sorted we can use binary search
func findKeyRotated(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + ((r - l) / 2)
		if nums[m] == target {
			return m
		} else if target >= nums[l] {
			if target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			if target > nums[r] {
				r = m - 1
			} else {
				l = m + 1
			}

		}
	}
	return -1
}

// a[i]=i  a[i]=-1, a[]!=i
// {-1, -1, 6, 1, 9, 3, 2, -1, 4, -1}
func Rearrange(a []int) {
	for i := 0; i < len(a); i++ {
		if a[i] != i && a[i] != -1 {
			t := a[i]
			for a[t] != -1 && a[t] != t {
				t2 := a[t]
				a[t] = t
				t = t2
			}
			a[t] = t
			if a[i] != i {
				a[i] = -1
			}

		}
	}
}

func BuyAndsell(s []int) int {
	n := len(s)
	maxProfit := 0
	maxPrice := s[n-1]
	for j := n - 2; j >= 0; j-- {
		if maxPrice < s[j] {
			maxPrice = s[j]
		}
		local := maxPrice - s[j]
		if maxProfit < local {
			maxProfit = local
		}
	}
	return maxProfit
}

func BuyAndsell2(prices []int) int {
	n := len(prices)
	firstBuySellProfits := make([]int, n)
	firstBuySellProfits[0] = 0
	minPriceSofar := prices[0]
	maxProfit := 0

	for i := 1; i < n; i++ {
		maxProfit = utils.MaxOf(prices[i]-minPriceSofar, maxProfit)
		minPriceSofar = utils.Min(minPriceSofar, prices[i])
		firstBuySellProfits[i] = maxProfit
	}

	maxPriceSofar := prices[n-1]

	for j := n - 2; j > 0; j-- {
		maxPriceSofar = utils.MaxOf(prices[j], maxPriceSofar)
		maxProfit = utils.MaxOf(maxProfit, maxPriceSofar-prices[j]+firstBuySellProfits[j-1])
	}
	return maxProfit
}

/*

	 {1, 2, 3, 4, 5, 6, 7}
k=3, {5, 6, 7, 1, 2, 3, 4}
k=2, {6, 7, 1, 2, 3, 4, 5}
k=1, {7, 1, 2, 3, 4, 5, 6}
*/

func Rotation(nums []int, k int) {
	l := 0
	r := len(nums) - 1 - k
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
	l = len(nums) - k
	r = len(nums) - 1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
	l = 0
	r = len(nums) - 1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}

}

// {2,3,1,1,4}
func Jump(nums []int) bool {
	lidx := len(nums) - 1
	if lidx == -1 {
		return false
	}
	if lidx == 0 {
		return true
	}

	temp := make([]bool, len(nums))
	END := lidx
	for i := lidx - 1; i >= 0; i-- {
		maxJupms := nums[i]
		for k := maxJupms; k > 0; k-- {
			if i+k <= END {
				if i+k == END || temp[i+k] {
					temp[i] = true
				}
			}
		}
	}
	return temp[0]
}

// Sorted array
func removeDuplicates(input []int) {
	i := 0
	for j := 1; j < len(input); j++ {
		if input[j] != input[i] {
			i = i + 1
			input[i] = input[j]
		}
	}
	i++
	for ; i < len(input); i++ {
		input[i] = -1
	}
}

// {1, 2, 0, 0, 0, 3, 6};
// Method-1 : whenever zero element move all next elements to before then place zero at the end.
// Method-2 : Counting zeros and moving next nonzero elemnt
func MovallZeros(a []int) {
	widx := -1
	for i := 0; i < len(a); i++ {
		if a[i] != 0 {
			widx += 1
			a[widx] = a[i]
		}
	}
	for widx += 1; widx < len(a); widx++ {
		a[widx] = 0
	}
}
