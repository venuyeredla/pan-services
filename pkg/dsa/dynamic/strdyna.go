package dynamic

import (
	"github.com/venuyeredla/pan-services/pkg/dsa/utils"
)

/*
saturday (8)
sunday (6)
*/
func ediDistance(s1, s2 string) int {
	if len(s1) == 0 {
		return len(s2)
	}
	if len(s2) == 0 {
		return len(s1)
	}
	if s1[0] == s2[0] {
		return ediDistance(s1[1:], s2[1:])
	}
	u := ediDistance(s1[1:], s2[1:])
	d := ediDistance(s1, s2[1:])
	i := ediDistance(s1[1:], s2)

	return 1 + utils.Minimum(d, u, i)
}

func ediDistanceD(s1, s2 string) int {

	y := len(s1)
	x := len(s2)
	sm := utils.GetMatrix(x+1, y+1)

	//First line fillling
	for k := 0; k <= y; k++ {
		sm[0][k] = k
	}
	//First column filling
	for k := 0; k <= x; k++ {
		sm[k][0] = k
	}

	for i := 1; i <= x; i++ {
		for j := 1; j <= y; j++ {
			if s2[i-1] == s1[j-1] {
				sm[i][j] = sm[i-1][j-1]
			} else {
				sm[i][j] = 1 + utils.Minimum(sm[i-1][j-1], sm[i-1][j], sm[i][j-1])
			}
		}
	}

	return sm[x][y]
}

/* Longest common subsequence
aggtab
gxtxayb
case:1 "" "gxtxayb" or  aggtab  ""

LCS : gtab => 4

Botto up approach

*/

func Lcs(s1, s2 string) int {
	if len(s1) == 0 || len(s2) == 0 {
		return 0
	}
	if s1[0] == s2[0] {
		return 1 + Lcs(s1[1:], s2[1:])
	} else {
		lcs := utils.MaxOf(Lcs(s1[1:], s2), Lcs(s1, s2[1:]))
		return lcs
	}
}

/*
	 Longest repeating subsequence
		aabebcdd => abd
		aabebcdd
		(7,7) => (0,0)
*/

func Lrs(s string, m, n int) int {
	if m == -1 || n == -1 {
		return 0
	}
	if s[m] == s[n] && m != n {
		return 1 + Lrs(s, m-1, n-1)
	} else {
		lcs := utils.MaxOf(Lrs(s, m-1, n), Lrs(s, m, n-1))
		return lcs
	}
}

func LcsD(s1, s2 string) int {

	y := len(s1)
	x := len(s2)
	sm := utils.GetMatrix(x+1, y+1)

	//First line fillling
	for k := 0; k <= y; k++ {
		sm[0][k] = 0
	}
	//First column filling
	for k := 0; k <= x; k++ {
		sm[k][0] = 0
	}

	for i := 1; i <= x; i++ {
		for j := 1; j <= y; j++ {
			if s2[i-1] == s1[j-1] {
				sm[i][j] = 1 + sm[i-1][j-1]
			} else {
				sm[i][j] = utils.MaxOf(sm[i-1][j], sm[i][j-1])
			}
		}
	}
	return sm[x][y]
}

/*
	Longest Palindromic subsequence
	bbabcbcab => babcbab


	s : bbabcbcab
	r : bacbcbabb


*/

func Lps(s string) int {
	return LpsR(s, 0, len(s)-1)
}

func LpsR(s string, l, r int) int {
	if l == r {
		return 1
	}
	if l+1 == r && s[l+1] == s[r] {
		return 2
	}

	if s[l] == s[r] {
		return LpsR(s, l+1, r-1) + 2
	}
	return utils.MaxOf(LpsR(s, l+1, r), LpsR(s, l, r-1))
}

/*
Longest Palindromic subsequence
bbabcbcab => babcbab

We can use : LCS of orignal and reversed strings.

	s : bbabcbcab
	r : bacbcbabb
*/
func LpsD(s string) int {

	x := len(s)
	sm := utils.GetMatrix(x, x)

	for i := 0; i < x; i++ {
		sm[i][i] = 1
	}
	for i := 1; i < x; i++ {
		for j := i + 1; j < x; j++ {
			if s[i] == s[j] {
				sm[i][j] = 2 + sm[i+1][j-1]
			} else if s[i] == s[j-1] {
				sm[i][j] = 2
			} else {
				sm[i][j] = utils.MaxOf(sm[i+1][j-1], sm[i+1][j-1])
			}
		}
	}
	return 0 // sm[x][y]
}

/*
geeks => g e, e, k , s , ee
*/
func palindromPartiions(s string) int {
	if len(s) == 0 || len(s) == 0 {
		return 0
	}
	return 1
}

/*
For with recursion.
*/
func wordBreak(words []string, s1 string) bool {
	if len(s1) == 0 {
		return false
	}

	return false
}

/*

Text = "baaabab",
Pattern = “*****ba*****ab", output : true
Pattern = "baaa?ab", output : true
Pattern = "ba*a?", output : true
*/

func wildcard(text, pattern string) bool {
	if len(pattern) == 0 {
		return false
	}

	return false
}
