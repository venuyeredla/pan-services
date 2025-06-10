package dynamic

import (
	"fmt"

	"github.com/venuyeredla/pan-services/pkg/dsa/utils"
)

/*
	Subset sum problem is to find subset of elements that are selected from a given set whose sum adds up to a given number K.
	Approach :
		1. Brute forece. Genereate all subsets. loop for subsetsum. Recursion.
		2. Dynamic programming

// S,  s={a  a S  }  Sum(s) =K

3,2,7,1  = Sum =6
*/
func SubsetSum(A []int, sum int) bool {
	if sum == 0 {
		return true
	}
	if sum < 0 {
		return false
	}
	if sum-A[0] == 0 {
		return true
	}
	if len(A) > 1 {
		including := SubsetSum(A[1:], sum-A[0])
		excluding := SubsetSum(A[1:], sum)

		return including || excluding
	}
	return false
}

func SubsetSumD(A []int, sum int) bool {
	rows := len(A)
	columns := sum + 1
	sm := utils.GetMatrix(rows, columns)
	for i := 0; i < len(A); i++ {
		sm[i][0] = 1
	}
	for i := 1; i < columns; i++ {
		if A[0] == i {
			sm[0][i] = 1
		} else {
			sm[0][i] = -1
		}
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			if A[i] == j || sm[i-1][j] == 1 || ((j-A[i]) >= 0 && sm[i-1][j-A[i]] == 1) {
				sm[i][j] = 1
			} else {
				sm[i][j] = -1
			}
		}
	}
	if sm[rows-1][columns-1] == 1 {
		return true
	} else {
		return false
	}

}

/*
Approach :
 1. Bruteforce : n*(n+1)/2  => SumRange(i,j) - Involves recompuations (Dynamic programming)
 2. Prefix sum or segement or BIT. O(n,2)
 3. Dynamic programming
*/
func LargestSumSubArray(a []int) {
	//StringJoiner stringJoiner=new StringJoiner(", ");
	utils.Printable(a, 0, len(a)-1)
	sum := 0
	fidx := 0
	toIdx := 0
	maxSum := 0
	for j := 0; j < len(a); j++ {
		sum = sum + a[j]
		if maxSum <= sum {
			maxSum = sum
			toIdx = j
		}
		if sum < 0 {
			sum = 0
			fidx = j
		}
		//stringJoiner.add(sum + "")
	}
	//System.out.print(stringJoiner.toString())
	fmt.Printf(" => {%v,%v = %v} \n", fidx+1, toIdx, maxSum)
}

/*
Input: arr[] = {3, 10, 2, 1, 20} => 3, 10, 20

	Approach :
	1. This is combination problem and for each elements try find lenght by including or excluding next element.
	Bottom up apporoach.
	1. len(A)==cidx return 0;
	2. if A[cidx] >= A[lidx]  return a+lis(A, cidx+1, cix)
	3. If not equal.
		a. Keep lidx as it is and increas cidx  -- Exclusion.
		b. Assuming this could other sequnce minimum element and start counting again.


	Approach - 2.

	Dynamic proramming.

Output: 3
*/
func longIncreasingSubseq(A []int, cidx, lastidx int) int {
	if len(A) == cidx { // one element case;
		return 0
	}
	if cidx == lastidx && len(A) == cidx+1 {
		return 1
	}
	includingCount := 0
	if A[cidx] >= A[lastidx] {
		return 1 + longIncreasingSubseq(A, cidx+1, cidx)
	} else {
		longIncreasingSubseq(A, cidx+1, cidx)
	}
	excludingCount := longIncreasingSubseq(A, cidx, cidx)
	// Max(exclusion, inclusion)
	return utils.MaxOf(includingCount, excludingCount)
}

/*
	 Input: arr[] = {3, 10, 2, 1, 20} => 3, 10, 20
					 {1, 2, 1, 1, 3}
*/
func LISD(A []int) int {
	solution := make([]int, len(A))
	lidx := 0
	solution[0] = 1
	for i := 1; i < len(A); i++ {
		if A[i] >= A[lidx] {
			solution[i] = 1 + solution[lidx]
			lidx = i
		} else {

		}
	}
	return A[len(A)-1]
}

/*
Input arr[] = {1, 11, 2, 10, 4, 5, 2, 1};
Output: 6 (A Longest Bitonic Subsequence of length 6 is 1, 2, 10, 4, 2, 1)

Increasing then decreasing
*/

func longBitonicSubseq(A []int) int {

	return 0
}
