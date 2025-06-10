package array

import (
	"testing"

	"github.com/venuyeredla/pan-services/pkg/dsa/utils"
)

func TestRemoveDuplicates(t *testing.T) {
	arr := []int{2, 3, 5, 5, 7, 11, 11, 11, 13}
	removeDuplicates(arr)
}

func TestMoveAllzeros(t *testing.T) {
	arr := []int{1, 0, 2, 0, 0, 3}
	expected := []int{1, 2, 3, 0, 0, 0}
	MovallZeros(arr)
	result, msg := utils.AssertEquals(expected, arr, false)
	if !result {
		t.Errorf(msg)
		t.Fail()
	}

}
