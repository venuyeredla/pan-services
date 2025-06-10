package text

import (
	"fmt"
	"testing"
)

func TestSuffixTree(t *testing.T) {
	//t.Skip()
	//PatMatchSuffix("venugopal", "gopal")
	result := PatMatchSuffixArray("venugopal", "gopal")
	if result != 4 {
		t.FailNow()
	} else {
		fmt.Printf("Pattern matched at index = %v", result)
	}
}
