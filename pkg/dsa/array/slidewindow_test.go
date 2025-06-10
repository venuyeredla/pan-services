package array

import "testing"

func TestShortestSubarray(t *testing.T) {
	input := []int{84, -37, 32, 40, 95}
	expected := 3
	result := shortestSubarray(input, 167)
	if result != expected {
		t.Errorf("Expected =%v, Acutal= %v", expected, result)
		t.FailNow()
	}
}
