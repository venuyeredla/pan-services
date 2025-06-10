package text

import (
	"fmt"
	"testing"
)

type TestIO struct {
	Input  string
	Output string
}

func TestReverse(t *testing.T) {
	io := make([]TestIO, 0, 4)
	io = append(io, TestIO{"", ""})
	io = append(io, TestIO{"v", "v"})
	io = append(io, TestIO{"ve", "ev"})
	io = append(io, TestIO{"venugopal", "lapogunev"})
	for _, val := range io {
		reverse := ""
		if reverse != val.Output {
			fmt.Println(val.Output, reverse)
			t.Fail()
		}
	}
}

func TestString(t *testing.T) {
	MultiStr("0033", "22")
}

func TestRToD(t *testing.T) {
	num := RomanDecimal("MCMXCIV")
	if num == 1994 {
		t.Log("Not a valid decimal")
		t.Fail()
	}
}

func TestFuck(t *testing.T) {
	nums := []byte{97, 97, 98, 98, 99, 99, 99}
	min, val := compress(nums)
	fmt.Printf("value =%v \n   an = %v\n", min, val)
}

func TestFuck2(t *testing.T) {
	val := 2
	fmt.Printf("value =%v \n   an = %v", val, val)
}
