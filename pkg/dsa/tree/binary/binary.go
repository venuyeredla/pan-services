package binary

import (
	"github.com/venuyeredla/pan-services/pkg/dsa/errors"
	"github.com/venuyeredla/pan-services/pkg/dsa/stack_queue"
	"github.com/venuyeredla/pan-services/pkg/dsa/tree"
	"github.com/venuyeredla/pan-services/pkg/dsa/types"
)

type COLOR uint
type TreeType uint

// TraversalType represents one of the three know traversals.
type TraversalType int

const (
	InOrder TraversalType = iota
	PreOrder
	PostOrder
)

const (
	LEVEL TreeType = iota
	AVL
	BST
	RB
)

const (
	RED COLOR = iota
	BLACK
)

type BinaryNode struct {
	key                 types.Hashable
	value               interface{}
	parent, left, right *BinaryNode
	height              int   //For avl tree
	color               COLOR // For redblack tree
}

type BinaryTree struct {
	root     *BinaryNode
	treeType TreeType
	count    int
}

func NewBinaryTree(tType TreeType) *BinaryTree {
	return &BinaryTree{treeType: tType}
}

func (self *BinaryTree) Put(key types.Hashable, value interface{}) (err error) {
	switch self.treeType {
	case LEVEL:
		self.root, _ = self.root.LevelPut(key, value) //// Average: O(log(n)) Worst: O(n)
		break

	case BST:
		self.root, _ = self.root.BstPut(key, value) //// Average: O(log(n)) Worst: O(n)
		break

	case AVL:
		self.root, _ = self.root.AvlPut(key, value)
		return nil

	case RB:
		self.root = RBPut(self.root, key)
		break

	default:
		break

	}
	return nil
}

func (self *BinaryTree) Get(key types.Hashable) (value interface{}, err error) {
	return self.root.Get(key)
}

func (self *BinaryNode) Get(key types.Hashable) (value interface{}, err error) {
	if self == nil {
		return nil, errors.NotFound(key)
	}
	if self.key.Equals(key) {
		return self.value, nil
	} else if key.Less(self.key) {
		return self.left.Get(key)
	} else {
		return self.right.Get(key)
	}
}

func (self *BinaryTree) Remove(key types.Hashable) (value interface{}, err error) {

	switch self.treeType {
	case RB:
		self.root = RBPut(self.root, key)
		break
	case AVL:
		new_root, value, err := self.root.AvlRemove(key)
		if err != nil {
			return nil, err
		}
		self.root = new_root
		return value, nil
	case LEVEL:
		RBPut(self.root, key)
		break
	}
	return value, nil
}

func (self *BinaryTree) Iterate() types.KVIterator {
	return self.root.Iterate()
}

func (self *BinaryTree) Items() (vi types.KIterator) {
	return types.MakeItemsIterator(self)
}

func (self *BinaryTree) Values() types.Iterator {
	return self.root.Values()
}

func (self *BinaryTree) Keys() types.KIterator {
	return self.root.Keys()
}

func (self *BinaryTree) Root() types.TreeNode {
	return self.root
}

func (self *BinaryTree) Size() int {
	return self.root.Size()
}

func (self *BinaryTree) Has(key types.Hashable) bool {
	return self.root.Has(key)
}

func (self *BinaryNode) Has(key types.Hashable) (has bool) {
	if self == nil {
		return false
	}
	if self.key.Equals(key) {
		return true
	} else if key.Less(self.key) {
		return self.left.Has(key)
	} else {
		return self.right.Has(key)
	}
}

// TreeNode interface
func (self *BinaryNode) Key() types.Hashable {
	return self.key
}

func (self *BinaryNode) Value() interface{} {
	return self.value
}

func (self *BinaryNode) GetChild(i int) types.TreeNode {
	return types.DoGetChild(self, i)
}

func (self *BinaryNode) ChildCount() int {
	return types.DoChildCount(self)
}

func (self *BinaryNode) Children() types.TreeNodeIterator {
	return types.MakeChildrenIterator(self)
}

// TreeNode interface

func (self *BinaryNode) Iterate() types.KVIterator {
	tni := tree.TraverseBinaryTreeInOrder(self)
	return types.MakeKVIteratorFromTreeNodeIterator(tni)
}

func (self *BinaryNode) Keys() types.KIterator {
	return types.MakeKeysIterator(self)
}

func (self *BinaryNode) Values() types.Iterator {
	return types.MakeValuesIterator(self)
}

// insert recusively adds a key+value in the tree.
// Lever order insertion using Queue
func (self *BinaryNode) LevelPut(key types.Hashable, value interface{}) (r *BinaryNode, updated bool) {

	//func insertInOrder(n *node, k int, v interface{}) (r *node, added bool) {
	newNode := &BinaryNode{key: key, value: value}
	if r = self; self == nil {
		r = newNode
		updated = true
		return r, updated
	}
	queue := new(stack_queue.Queue)
	queue.Init()
	queue.Push(self)

	for !queue.IsEmpty() {
		current := ToNode(queue.Pop())
		if current.left == nil {
			current.left = newNode
			break
		} else {
			queue.Push(current.left)
		}
		if current.right == nil {
			current.right = newNode
			break
		} else {
			queue.Push(current.right)
		}
	}

	return r, true
}

func ToNode(n1 interface{}) (n *BinaryNode) {
	//node(n)
	return
}

func (self *BinaryNode) pop_node(node *BinaryNode) *BinaryNode {
	if node == nil {
		panic("node can't be nil")
	} else if node.left != nil && node.right != nil {
		panic("node must not have both left and right")
	}

	if self == nil {
		return nil
	} else if self == node {
		var n *BinaryNode
		if node.left != nil {
			n = node.left
		} else if node.right != nil {
			n = node.right
		} else {
			n = nil
		}
		node.left = nil
		node.right = nil
		return n
	}

	if node.key.Less(self.key) {
		self.left = self.left.pop_node(node)
	} else {
		self.right = self.right.pop_node(node)
	}

	self.height = max(self.left.Height(), self.right.Height()) + 1
	return self
}

func (self *BinaryNode) push_node(node *BinaryNode) *BinaryNode {
	if node == nil {
		panic("node can't be nil")
	} else if node.left != nil || node.right != nil {
		panic("node now be a leaf")
	}

	if self == nil {
		node.height = 1
		return node
	} else if node.key.Less(self.key) {
		self.left = self.left.push_node(node)
	} else {
		self.right = self.right.push_node(node)
	}
	self.height = max(self.left.Height(), self.right.Height()) + 1
	return self
}

func (self *BinaryNode) Height() int {
	if self == nil {
		return 0
	}
	return self.height
}

func (self *BinaryNode) Size() int {
	if self == nil {
		return 0
	}
	return 1 + self.left.Size() + self.right.Size()
}

func (self *BinaryNode) Left() types.BinaryTreeNode {
	if self.left == nil {
		return nil
	}
	return self.left
}

func (self *BinaryNode) Right() types.BinaryTreeNode {
	if self.right == nil {
		return nil
	}
	return self.right
}

func (self *BinaryNode) _md(side func(*BinaryNode) *BinaryNode) *BinaryNode {
	if self == nil {
		return nil
	} else if side(self) != nil {
		return side(self)._md(side)
	} else {
		return self
	}
}

func (self *BinaryNode) lmd() *BinaryNode {
	return self._md(func(node *BinaryNode) *BinaryNode { return node.left })
}

func (self *BinaryNode) rmd() *BinaryNode {
	return self._md(func(node *BinaryNode) *BinaryNode { return node.right })
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/* RB tree insertion and deleetion */
/*
Properties.
1. Every nodes is either red or black
2. Root is always Black
3. A red node doesn't have red child/parent
3. All nil nodes considered as leafs and need to be blacks
4. Every path from a given node to any of it's descendents nil(leaves) has same no of black nodes.
*/

// While inserting a key we use two tools for balancing tree.
// 1. Recoloring and rotation
func RBPut(n *BinaryNode, key types.Hashable) (root *BinaryNode) {

	if n == nil {
		return &BinaryNode{key: key}
	} else if key.Less(n.key) {
		n.left = RBPut(n.left, key)
		n.left.parent = n

	} else {
		n.right = RBPut(n.right, key)
		n.right.parent = n
	}

	return nil
}

func RBDelete(n *BinaryNode) {

}

/* RB tree insertion and deleetion */

/* BST tree insertion and deleetion */

// Insert adds a given key+value to the tree and returns true if it was added.
// Average: O(log(n)) Worst: O(n)

func (t *BinaryTree) BstPut(k types.Hashable, v interface{}) (updated bool) {
	t.root, updated = t.root.BstPut(k, v)
	if updated {
		t.count++
	}
	return updated
}

// insert recusively adds a key+value in the tree.
func (n *BinaryNode) BstPut(k types.Hashable, v interface{}) (r *BinaryNode, added bool) {
	if r = n; n == nil {
		// keep track of how many elements we have in the tree
		// to optimize the channel length during traversal
		r = &BinaryNode{key: k, value: v}
		added = true
	} else if k.Less(n.key) {
		r.left, added = n.left.BstPut(k, v)
	} else if n.key.Less(k) {
		r.right, added = n.right.BstPut(k, v)
	}

	return
}

// Delete removes a given key from the tree and returns true if it was removed.
// Average: O(log(n)) Worst: O(n)
func (t *BinaryTree) Delete(k types.Hashable) (deleted bool) {
	n, deleted := delete(t.root, k)
	if deleted {
		// Handling the case of root deletion.
		if t.root.key.Equals(k) {
			t.root = n
		}
		t.count--
	}

	return deleted
}

// delete recursively deletes a key from the tree.
func delete(n *BinaryNode, k types.Hashable) (r *BinaryNode, deleted bool) {
	if r = n; n == nil {
		return nil, false
	}

	if k.Less(n.key) {
		r.left, deleted = delete(n.left, k)
	} else if n.key.Less(k) {
		r.right, deleted = delete(n.right, k)
	} else {
		if n.left != nil && n.right != nil {
			// find the right most element in the left subtree
			s := n.left
			for s.right != nil {
				s = s.right
			}
			r.key = s.key
			r.value = s.value
			r.left, deleted = delete(s, s.key)
		} else if n.left != nil {
			r = n.left
			deleted = true
		} else if n.right != nil {
			r = n.right
			deleted = true
		} else {
			r = nil
			deleted = true
		}
	}

	return
}

// Find returns the value found at the given key.
// Average: O(log(n)) Worst: O(n)
func (t *BinaryTree) Find(k types.Hashable) interface{} {
	return find(t.root, k)
}

func find(n *BinaryNode, k types.Hashable) interface{} {
	if n == nil {
		return nil
	}
	if n.key.Equals(k) {
		return n.value
	} else if k.Less(n.key) {
		return find(n.left, k)
	} else if n.key.Less(k) {
		return find(n.right, k)
	}

	return nil
}

// Clear removes all the nodes from the tree.
// O(n)
func (t *BinaryTree) Clear() {
	t.root = clear(t.root)
	t.count = 0
}

// clear recursively removes all the nodes.
func clear(n *BinaryNode) *BinaryNode {
	if n != nil {
		n.left = clear(n.left)
		n.right = clear(n.right)
	}
	n = nil

	return n
}

// Traverse provides an iterator over the tree.
// O(n)
func (t *BinaryTree) Traverse(tt TraversalType) <-chan interface{} {
	c := make(chan interface{}, t.count)
	go func() {
		switch tt {
		case InOrder:
			inOrder(t.root, c)
			break
		case PreOrder:
			preOrder(t.root, c)
			break
		case PostOrder:
			postOrder(t.root, c)
			break
		}
		close(c)
	}()
	return c
}

// inOrder returns the left, parent, right nodes.
func inOrder(n *BinaryNode, c chan interface{}) {
	if n == nil {
		return
	}

	inOrder(n.left, c)
	c <- n.value
	inOrder(n.right, c)
}

// preOrder returns the parent, left, right nodes.
func preOrder(n *BinaryNode, c chan interface{}) {
	if n == nil {
		return
	}

	c <- n.value
	preOrder(n.left, c)
	preOrder(n.right, c)
}

// postOrder returns the left, right, parent nodes.
func postOrder(n *BinaryNode, c chan interface{}) {
	if n == nil {
		return
	}
	postOrder(n.left, c)
	postOrder(n.right, c)
	c <- n.value
}
