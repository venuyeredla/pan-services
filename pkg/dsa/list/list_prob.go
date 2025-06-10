package list

import (
	"container/list"
	"fmt"
	"strings"

	"github.com/venuyeredla/pan-services/pkg/dsa/stack_queue"
)

/*
# Notes
Use sentinel(dummy head)

# Problems.
1. Merge two sorted lists.
2. Reverse SLL & DLL.
3. Cycles in SLL.
4. Test for overlappeing lists.
5. Remove duplicates from sorted list.
6. Cyclic right shift
7. Test for palindrome of list.
8. Add list based integers.
*/

type LNode struct {
	Value              any
	Prev, Next, Random *LNode
}

func goList() {
	list := list.New()
	list.PushBack(10)
	list.PushBack(12)

	list.PushFront(24)

	temp := list.Front()

	for temp != nil {
		fmt.Printf("%v  ", temp.Value)
		temp = temp.Next()
	}

	//list.Back()
	//list.Front()
	//list.PushBack()

}

/*
Trying to merge into first tree.
*/
func AddTwoNumbers(list1, list2 *list.List) *list.List {
	c1 := list1.Front()
	c2 := list2.Front()
	carry := 0
	for c1 != nil && c2 != nil {
		num1 := c1.Value.(int)
		num2 := c2.Value.(int)
		sum := num1 + num2 + carry
		if sum > 9 {
			diggit := sum % 10
			carry = sum / 10
			c1.Value = diggit
		} else {
			c1.Value = sum
		}
		c1 = c1.Next()
		c2 = c2.Next()
	}
	return list1
}

/*
1. Iteration
2. Stack
3. Recursion
*/
func listReversalStack(h *LNode) *LNode {

	pointer := h
	stack := stack_queue.NewStack(10)
	for pointer != nil {
		temp := pointer
		pointer = pointer.Next
		temp.Next = nil
		stack.Push(temp)
	}
	var newHead, tail *LNode
	for !stack.IsEmpty() {
		newNode := stack.Pop().(*LNode)
		if newHead == nil {
			newHead, tail = newNode, newNode
		} else {
			tail.Next = newNode
			tail = newNode
		}
	}
	return newHead
}

func listReversalRecursion(h *LNode) *LNode {
	if h == nil || h.Next == nil {
		return h
	}
	temp := h
	h = h.Next
	temp.Next = nil
	rev := listReversalRecursion(h)
	tail := rev
	for tail != nil {
		tail = tail.Next
	}
	tail.Next = temp
	return rev
}

func deepClone(head *LNode) *LNode {
	iter := head
	nodeMapping := make(map[*LNode]*LNode)

	for iter != nil {
		nodeMapping[iter] = &LNode{Value: iter.Value}
		iter = iter.Next
	}
	iter = head
	for iter != nil {
		nodeMapping[iter].Next = nodeMapping[iter.Next]
		nodeMapping[iter].Random = nodeMapping[iter.Random]
		iter = iter.Next
	}
	return nodeMapping[head]
}

func reorganizeString(s string) string {
	var h, t *LNode
	for i := range s {
		nn := &LNode{Value: s[i : i+1]}
		if h == nil {
			h = nn
			t = nn
		} else {
			t.Next = nn
			t = t.Next
		}
	}
	temp := h
	for temp != nil && temp.Next != nil {
		if temp.Next.Value.(string) == temp.Value.(string) {
			nt := temp.Next
			if nt.Next == nil {
				return ""
			}
			temp.Next = nt.Next
			nt.Next = nil
			t.Next = nt
			t = t.Next
		} else {
			temp = temp.Next
		}
	}
	var sb strings.Builder
	temp = h
	for temp != nil {
		sb.WriteString(temp.Value.(string))
		temp = temp.Next
	}
	return sb.String()
}
