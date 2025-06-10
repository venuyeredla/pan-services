package text

import (
	"fmt"
	"strings"
)

func compress(chars []byte) (int, string) {
	if len(chars) == 0 {
		return 0, ""
	}
	var sb strings.Builder
	prev := chars[0]
	prevCount := 1
	//widx := 0
	for i := 1; i < len(chars); i++ {
		if prev == chars[i] {
			prevCount++
		} else {
			sb.WriteString(string(prev))
			sb.WriteString(fmt.Sprint(prevCount))
			prev = chars[i]
			prevCount = 1
		}
	}
	sb.WriteByte(prev)
	sb.WriteString(fmt.Sprint(prevCount))
	chars = make([]byte, sb.Len())
	copy(chars, []byte(sb.String()))
	return sb.Len(), sb.String()
}
