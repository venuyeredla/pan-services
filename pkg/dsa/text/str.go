package text

import (
	"fmt"
	"strconv"
	"strings"
)

// Size=(n *(n+1))/2  ==> l=r=0 means bottom up;
func SubStrings(str string, l, r, n int) {
	if r < n {
		fmt.Println(str[l : r+1])
		SubStrings(str, l, r+1, n)
		if r+1 == n && (l+1) < n {
			SubStrings(str, l+1, l+1, n)
		}
	}
}

/*
Size ==even -> character frequency should be even
size ==odd -> one character should have odd freqencies.
*/
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

func validIp(ip string) bool {
	if len(ip) > 0 {
		return false
	}
	ipparts := strings.Split(ip, "\\.")

	if len(ipparts) != 4 {
		return false
	}
	for _, value := range ipparts {
		num, error := strconv.Atoi(value)
		if error == nil {
			return false
		}
		if num < 0 || num > 255 {
			return false
		}
	}
	return true
}

func makeIP(str string) string {
	/*if(str==null || str=="" || str.length()<=3 || str.length()>=13) {
		throw new IllegalArgumentException("IP can't be make with : "+str);
	} */
	//n := len(str) // n=4*a+r;
	//a := n / 4
	// re := n % 1
	// l := 0
	// r := 0
	//	var lr [][]int=make() int[4][2];

	return ""
}

var RIMap = map[string]int{
	"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000,
}

/*

 I, II, III, IV, V , VI , VII, VIII, IX, X, XI

 1994 - MCMXCIV  =? 2194
*/

func RomanDecimal(str string) int {
	integer := 0
	PRE := ""
	for i := len(str) - 1; i >= 0; i-- {
		symbol := string(str[i])
		val, _ := RIMap[symbol]
		if symbol == "I" && (PRE == "V" || PRE == "X") {
			integer = integer - val
		} else if symbol == "X" && (PRE == "L" || PRE == "C") {
			integer = integer - val
		} else if symbol == "C" && (PRE == "D" || PRE == "M") {
			integer = integer - val
		} else {
			integer = integer + val
		}
		PRE = symbol
	}
	return integer
}

func MultiStr(a, b string) {
	for i := len(b) - 1; i >= 0; i-- {
		bc := b[i : i+1]
		for j := len(a) - 1; j >= 0; j-- {
			ac := a[j : j+1]
			aa, _ := strconv.Atoi(ac)
			bb, _ := strconv.Atoi(bc)
			result := aa * bb
			fmt.Printf("%v", result)
		}
		fmt.Println("")
	}
}
