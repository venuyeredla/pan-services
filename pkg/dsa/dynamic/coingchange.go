package dynamic

import "fmt"

/*
Recuriosn {1,2,3} -? 4

	How do you make if you have one coin.
*/
func CoinChange(A []int, S int) int {

	if S == 0 {
		fmt.Printf("One pattern completed")
		return 1
	}
	if S < 0 {
		return 0
	}
	fmt.Printf("Requried %v Adding =%v \n", S, A[0])
	count1 := CoinChange(A, S-A[0])
	count2 := 0
	if len(A) > 1 {
		count2 = CoinChange(A[1:], S)
	}

	return count1 + count2
}
