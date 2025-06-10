package binary

import (
	"fmt"
	"testing"

	"github.com/venuyeredla/pan-services/pkg/dsa/types"
)

func createTree(arr []int) *BinaryTree {
	tree := NewBST()
	for _, val := range arr {
		tree.BstPut(types.Int(val), types.Int(val))
	}
	c := tree.Traverse(InOrder)
	for ele := range c {
		fmt.Printf("%v ,", ele)
	}
	return tree
}

func TestKthsmall(t *testing.T) {
	input := []int{3, 1, 2, 4}
	tree := createTree(input)
	value := KthNodeValue(tree.root, 1)

	fmt.Printf("Output : %v", value)

}
