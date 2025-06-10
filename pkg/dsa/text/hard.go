package text

import (
	"strings"
)

func TextJustify(words []string, maxWidth int) []string {
	solution := make([]string, 0)
	sb := make([]string, 0)
	available := maxWidth
	for _, word := range words {
		if available < len(word) {
			solution = append(solution, justify(sb, available, false))
			available = maxWidth
			sb = make([]string, 0)
		}
		sb = append(sb, word)
		available = available - len(word) - 1
	}
	if len(sb) > 0 {
		solution = append(solution, justify(sb, available, true))
	}
	return solution
}

func justify(strs []string, scount int, isLast bool) string {
	scount = scount + len(strs)
	spaces := make([]int, len(strs))
	if !isLast {
		j := 0
		for ; scount > 0; scount-- {
			spaces[j]++
			j++
			if j >= len(spaces)-1 {
				j = 0
			}

		}
	} else {
		i := 0
		for ; i < len(spaces)-1; i++ {
			spaces[i] = 1
			scount--
		}
		spaces[i] = scount
	}

	var sb strings.Builder
	for i, ws := range strs {
		sb.WriteString(ws)
		for i := spaces[i]; i > 0; i-- {
			sb.WriteString("*")
		}

	}
	return sb.String()
}

/*
Approach :Sliding window.

1. Take substring of total lengths every time

func findSubstring(s string, words []string) []int {
	if len(words) == 0 || len(s) == 0 {
		return []int{}
	}
	wl := len(words[0])
	//subLength := len(words) * wl

	var hash = func(s string) int {
		hash := 0
		for _, val := range s {
			hash += int(val)
		}
		return hash
	}
	strHash := 0
	hashSet := make(map[int]bool)
	for _, word := range words {
		temp := hash(word)
		hashSet[temp] = true
		strHash += temp
	}

	var isValid = func(str string) bool {
		if strHash == hash(str) {
			for i := 0; i <= len(str)-wl; i++ {
				w := str[i : i+wl]
				if _, exist := hashSet[hash(w)]; !exist {
					return false
				}
			}
			return true
		}
		return false
	}
	answer := make([]int, 0)
	/* for i := 0; i <= len(s)-subLength; i = i + 0 {
		if isValid(s[i : i+subLength])

		answer = append(answer, i)
			i = i + wl
		} else {
			i++
		}
	}
	return answer
}

*/
