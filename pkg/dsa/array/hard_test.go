package array

import (
	"testing"

	"github.com/venuyeredla/pan-services/pkg/dsa/utils"
)

func TestPrdoductExceptItself(t *testing.T) {
	input := []int{-1, 1, 0, -3, 3} // 1, 0, 2
	expected := []int{0, 0, 9, 0, 0}
	output := productExceptSelf(input)
	utils.AssertEquals(expected, output, false)
}

func TestDuplicate(t *testing.T) {
	input := []int{1, 3, 4, 2, 2}
	expected := 2
	output := findDuplicate(input)
	if output != expected {
		t.Errorf("Expected = %v and Actual =%v", expected, output)
		t.FailNow()
	}
}

func TestMinJumps(t *testing.T) {
	input := []int{1, 3, 5, 8, 9, 2, 6, 7, 6, 8, 9}
	expected := 3
	output := minJumpsRequred(input)
	if output != expected {
		t.Errorf("Expected = %v and Actual =%v", expected, output)
		t.FailNow()
	}
}
