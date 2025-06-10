package compress

import (
	"fmt"
	"testing"
)

func TestAbsCompress(t *testing.T) {
	fmt.Println("Testing ABS compression")
	AbsCompress("00101000101110101001010")
}
