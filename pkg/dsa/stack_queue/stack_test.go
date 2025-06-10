package stack_queue

import (
	"fmt"
	"testing"

	"github.com/venuyeredla/pan-services/pkg/dsa/utils"
)

func TestSPushPop(t *testing.T) {
	fmt.Println("Running push operation")
	s := NewStack(iterations / 2)
	for i := 0; i < iterations; i++ {
		s.Push(i)
	}
	for i := iterations - 1; i >= 0; i-- {
		testPopS(t, s, i)
	}
}

func TestInitPushSmallestStack(t *testing.T) {
	// Arrange.
	s := NewStack(1)
	for i := 0; i < 4; i++ {
		s.Push(i)
	}
	// Assert.
	for i := 3; i >= 0; i-- {
		testPopS(t, s, i)
	}
}

func TestPeek(t *testing.T) {
	s := NewStack(10)
	s.Push("a")
	testPeek(t, s, "a")

	s.Push("b")
	testPeek(t, s, "b")

	s.Pop()
	testPeek(t, s, "a")

	s.Pop()
	testPeek(t, s, nil)
}

func TestSLen(t *testing.T) {
	s := NewStack(iterations / 4)
	for i := 0; i < iterations; i++ {
		s.Push(i)
	}
	if l := s.Len(); l != iterations {
		t.Errorf("Stack length was expected to be %v, but is %v", iterations, l)
	}
	s.Pop()
	if l := s.Len(); l != iterations-1 {
		t.Errorf("Stack length was expected to be %v, but is %v", iterations-1, l)
	}
}

func TestSIsEmpty(t *testing.T) {
	s := NewStack(2)

	if s.IsEmpty() != true {
		t.Errorf("Stack should be empty")
	}

	s.Push(1)

	if s.IsEmpty() != false {
		t.Errorf("Stack should not be empty")
	}
}

func testPopS(t *testing.T, s *Stack, e interface{}) {
	if v := s.Pop(); v != e {
		t.Errorf("Popping expected %v, got %v", e, v)
	}
}

func testPeek(t *testing.T, s *Stack, e interface{}) {
	if v := s.Peek(); v != e {
		t.Errorf("Peeking expected %v, got %v", e, v)
	}
}

func BenchmarkPushNoResize(b *testing.B) {
	s := NewStack(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}

func BenchmarkPushResize(b *testing.B) {
	s := NewStack(b.N / 2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}

func BenchmarkPopS(b *testing.B) {
	s := NewStack(b.N)
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}

/* stack problems test */
func TestItoP(t *testing.T) {
	//ItoP("a+b*c+d")
	inputs := []string{"a+b*c+d", "((a+b)-c*(d/e))+f"}
	outputs := []string{"abc*+d+", "ab+cde/*-f+"}
	for idx, input := range inputs {
		output := InfixToPostfix(input)
		if outputs[idx] != output {
			t.Errorf("Input =%v , Expected= %v   , Output= %v", input, outputs[idx], output)
		}
	}
}

func TestCalculator(t *testing.T) {
	inputs := []string{"(1+(4+5+2)-3)+(6+8)"}
	outputs := []int{23}
	for i := 0; i < len(inputs); i++ {
		result := calculate(inputs[i])
		if result != outputs[i] {
			fmt.Printf("Input, exipectd,   actual = %v , %v, %v  ", inputs[i], outputs[i], result)
			t.FailNow()
		}
	}
}

func TestBalanced(t *testing.T) {
	balanced := isBalanced("[()]{}{[()()]()}")
	if !balanced {
		t.Fail()
	}
}

func TestStackSpan(t *testing.T) {
	input := []int{100, 80, 60, 70, 60, 75, 85} // 1, 0, 2
	expected := []int{1, 1, 1, 2, 1, 4, 6}
	output := stackSpan(input)
	success, message := utils.AssertEquals(expected, output, false)
	if !success {
		fmt.Println(message)
	}
}

func TestLongestParenthesis(t *testing.T) {
	inputs := []string{"(()", ")()())", "(())"}
	outputs := []int{2, 4, 4}
	for i := 0; i < len(inputs); i++ {
		result := longestValidParentheses(inputs[i])
		if result != outputs[i] {
			fmt.Printf("Input, exipectd,   actual = %v , %v, %v  ", inputs[i], outputs[i], result)
			t.FailNow()
		}
	}
}
