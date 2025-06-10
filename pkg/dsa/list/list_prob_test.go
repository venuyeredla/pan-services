package list

import (
	"container/list"
	"fmt"
	"testing"
)

func TestInbuiltList(t *testing.T) {
	goList()
}

func TestAddTwo(t *testing.T) {

	list1 := list.New()
	list2 := list.New()
	input := []int{2, 4, 3}
	for _, val := range input {
		list1.PushBack(val)
	}

	input2 := []int{5, 6, 4}
	for _, val := range input2 {
		list2.PushBack(val)
	}

	result := AddTwoNumbers(list1, list2)
	temp := result.Front()
	for temp != nil {
		fmt.Print(temp.Value.(int))
		temp = temp.Next()
	}

}

func TestReorganizeString(t *testing.T) {
	inputs := []string{"aab", "aaab"}
	outputs := []string{"aba", ""}
	for i := 0; i < len(inputs); i++ {
		result := reorganizeString(inputs[i])
		if result != outputs[i] {
			fmt.Printf("Input, exipectd,   actual = %v , %v, %v  ", inputs[i], outputs[i], result)
			t.FailNow()
		}
	}
}
