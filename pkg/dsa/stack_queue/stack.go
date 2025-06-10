package stack_queue

import (
	"strconv"
	"strings"

	"github.com/venuyeredla/pan-services/pkg/dsa/utils"
)

/*

   stack:=make([]string,0,len(s))
   var peek=func() string { return stack[len(stack)-1]}
   var push=func(ele string){ stack = append(stack, ele)}
   var pop=func () string{
       popped :=stack[len(stack)-1]
       stack=stack[:len(stack)-1]
        return popped
   }

*/

type Stack struct {
	storage []any
	size    int
}

func NewStack(size int) *Stack {
	if size == 0 {
		size = 1
	}
	stack := &Stack{storage: make([]any, size), size: -1}
	return stack
}

func (stack *Stack) Push(v any) {
	if stack.size+1 == cap(stack.storage) {
		ns := make([]any, cap(stack.storage)*2)
		copy(ns, stack.storage)
		stack.storage = ns
	}
	stack.size++
	stack.storage[stack.size] = v
}

func (stack *Stack) Pop() any {
	if stack.size < 0 {
		return nil
	}
	v := stack.storage[stack.size]
	stack.storage[stack.size] = nil
	stack.size--
	return v
}

func (stack *Stack) Peek() any {
	if stack.size < 0 {
		return nil
	}
	return stack.storage[stack.size]
}

func (stack *Stack) IsEmpty() bool {
	return stack.Len() == 0
}

func (stack *Stack) Len() int {
	return stack.size + 1
}

/*   Stack problems starting  */

// a+b*c+c -> abc*+d+
// a+b*(c^d-e)^(f+g*h)-i -> abcd^e-fgh*+^*+i-
func InfixToPostfix(expr string) string {
	preceMap := map[string]int{
		"+": 1, "-": 1,
		"%": 2, "/": 2, "*": 2,
		"^": 3,
	}
	stack := NewStack(len(expr))
	var precedence = func(s string) int {
		value, exist := preceMap[s]
		if exist {
			return value
		} else {
			return -1
		}
	}

	var isOperand = func(s string) bool {
		if s == "(" || s == ")" {
			return false
		}
		_, exist := preceMap[s]
		return !exist
	}

	var peek = func() string { return stack.Peek().(string) }
	var pop = func() string { return stack.Pop().(string) }
	var sb strings.Builder
	for _, c := range expr {
		s := string(c)
		//fmt.Print(s)
		if isOperand(s) {
			sb.WriteString(s)
		} else if s == "(" {
			stack.Push(s)
		} else if s == ")" {
			for !stack.IsEmpty() {
				t := peek()
				if t == "(" {
					stack.Pop()
					break
				}
				sb.WriteString(stack.Pop().(string))
			}
		} else {
			for !stack.IsEmpty() && (precedence(s) <= precedence(peek())) {
				sb.WriteString(pop())
			}
			stack.Push(s)
		}
	}
	for !stack.IsEmpty() {
		sb.WriteString(pop())
	}
	return sb.String()
}

func calculate(s string) int {
	operators := map[string]bool{"+": true, "-": true, "*": true, "/": true}
	s = InfixToPostfix(s)
	stack := NewStack(len(s))
	for i := 0; i < len(s); i++ {
		char := s[i : i+1]
		if _, isoperator := operators[char]; isoperator {
			b, _ := stack.Pop().(int)
			a, _ := stack.Pop().(int)
			switch char {
			case "+":
				stack.Push(a + b)
			case "-":
				stack.Push(a - b)
			case "*":
				stack.Push(a * b)
			case "/":
				stack.Push(a * b)
			}
		} else {
			val, error := strconv.Atoi(char)
			if error == nil {
				stack.Push(val)
			}
		}
	}
	return stack.Pop().(int)
}

func isBalanced(expr string) bool {
	stack := NewStack(len(expr))
	charMap := map[string]string{"}": "{", ")": "(", "]": "["}
	for i := range expr {
		s := expr[i : i+1]
		if val, exist := charMap[s]; exist && !stack.IsEmpty() {
			if stack.Pop().(string) != val {
				return false
			}
		} else {
			stack.Push(s)
		}
	}
	return stack.IsEmpty()
}

func IsWellFormedJson(expr string) bool {
	stack := NewStack(len(expr))
	charMap := map[string]string{"}": "{", "]": "[", "\"": "\"", ",": ",", ":": ":"}
	notJson := false
	for _, c := range expr {
		s := string(c)
		if _, exist := charMap[s]; exist {
			switch s {
			case "}":
				svalue, _ := stack.Pop().(string)
				if !(svalue == "}") {
					notJson = true
				}
			case "]":
				svalue, _ := stack.Pop().(string)
				if !(svalue == "[") {
					notJson = true
				}
			case "\"":

				break
			case ":":
				break
			case ",":
				break
			default:
			}

		} else {
			peek_value := stack.Pop().(string)
			if _, exist := charMap[peek_value]; !exist {
				stack.Push(s)
			}
		}
		if notJson {
			return notJson
		}
	}
	return stack.IsEmpty()
}

func stackSpan(prices []int) []int {
	spans := make([]int, len(prices))
	for i := 0; i < len(prices); i++ {
		span := 1
		j := i - 1
		for j >= 0 && prices[i] > prices[j] { // To reduce this loop we can use stack but depend on
			span++
			j--
		}
		spans[i] = span
	}
	return spans
}

// Total number of substrings n*(n+1)/2
func longestValidParentheses(s string) int {
	if len(s) < 1 {
		return 0
	}
	answer := 0
	stack := NewStack(len(s))
	for i := 0; i < len(s); i++ {
		char := s[i : i+1]
		if char == "(" {
			stack.Push(i)
		} else {
			stack.Pop()
			if stack.IsEmpty() {
				stack.Push(i)
			} else {
				answer = utils.MaxOf(answer, i-stack.Peek().(int))
			}
		}
	}
	return answer
}
