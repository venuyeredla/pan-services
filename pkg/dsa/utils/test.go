package utils

import (
	"fmt"
	"runtime/debug"
)

func (t *T) Assert(ok bool, msg string, vars ...interface{}) {
	if !ok {
		t.Log("\n" + string(debug.Stack()))
		t.Fatalf(msg, vars...)
	}
}

func (t *T) AssertNil(errors ...error) {
	any := false
	for _, err := range errors {
		if err != nil {
			any = true
			t.Log("\n" + string(debug.Stack()))
			t.Error(err)
		}
	}
	if any {
		t.Fatal("assert failed")
	}
}

func AssertEquals(expected, actual []int, log bool) (result bool, message string) {
	if log {
		Printable(expected, 0, len(expected)-1)
		Printable(actual, 0, len(actual)-1)
	}
	if len(expected) == len(actual) {
		for i := 0; i < len(expected); i++ {
			if expected[i] != actual[i] {
				return false, "Failed at index - " + fmt.Sprint(i) + " Expected = " + fmt.Sprint(expected[i]) + " Actual = " + fmt.Sprint(actual[i])
			}
		}
		return true, ""
	} else {
		return false, "Array lengths are unequal"
	}
}
