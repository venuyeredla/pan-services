package binary

import (
	"github.com/venuyeredla/pan-services/pkg/dsa/types"

	"github.com/venuyeredla/pan-services/pkg/dsa/errors"
)

/* AVL tree insertion and deleetion */

func NewBST() *BinaryTree {
	return &BinaryTree{treeType: BST}
}

/* AVL tree insertion and deleetion */

func (self *BinaryNode) AvlPut(key types.Hashable, value interface{}) (_ *BinaryNode, updated bool) {
	if self == nil {
		return &BinaryNode{key: key, value: value, height: 1}, false
	}

	if self.key.Equals(key) {
		self.value = value
		return self, true
	}

	if key.Less(self.key) {
		self.left, updated = self.left.AvlPut(key, value)
	} else {
		self.right, updated = self.right.AvlPut(key, value)
	}
	if !updated {
		self.height += 1
		return self.balance(), updated
	}
	return self, updated
}

func (self *BinaryNode) AvlRemove(key types.Hashable) (_ *BinaryNode, value interface{}, err error) {
	if self == nil {
		return nil, nil, errors.NotFound(key)
	}

	if self.key.Equals(key) {
		if self.left != nil && self.right != nil {
			if self.left.Size() < self.right.Size() {
				lmd := self.right.lmd()
				lmd.left = self.left
				return self.right, self.value, nil
			} else {
				rmd := self.left.rmd()
				rmd.right = self.right
				return self.left, self.value, nil
			}
		} else if self.left == nil {
			return self.right, self.value, nil
		} else if self.right == nil {
			return self.left, self.value, nil
		} else {
			return nil, self.value, nil
		}
	}
	if key.Less(self.key) {
		self.left, value, err = self.left.AvlRemove(key)
	} else {
		self.right, value, err = self.right.AvlRemove(key)
	}
	if err != nil {
		return self.balance(), value, err
	}
	return self, value, err
}

func (self *BinaryNode) rotate_right() *BinaryNode {
	if self == nil {
		return self
	}
	if self.left == nil {
		return self
	}
	new_root := self.left.rmd()
	self = self.pop_node(new_root)
	new_root.left = self.left
	new_root.right = self.right
	self.left = nil
	self.right = nil
	return new_root.push_node(self)
}

func (self *BinaryNode) rotate_left() *BinaryNode {
	if self == nil {
		return self
	}
	if self.right == nil {
		return self
	}
	new_root := self.right.lmd()
	self = self.pop_node(new_root)
	new_root.left = self.left
	new_root.right = self.right
	self.left = nil
	self.right = nil
	return new_root.push_node(self)
}

func (self *BinaryNode) balance() *BinaryNode {
	if self == nil {
		return self
	}
	for abs(self.left.Height()-self.right.Height()) > 2 {
		if self.left.Height() > self.right.Height() {
			self = self.rotate_right()
		} else {
			self = self.rotate_left()
		}
	}
	return self
}
