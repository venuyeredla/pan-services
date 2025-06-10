package binary

import (
	crand "crypto/rand"
	eb "encoding/binary"
	"fmt"
	mrand "math/rand"
	"testing"

	"github.com/venuyeredla/pan-services/pkg/dsa/utils"

	"github.com/venuyeredla/pan-services/pkg/dsa/types"
)

func TestInsert(t *testing.T) {
	expected := []int{5, 3, 7, 4, 6}
	bst := NewBST()

	for _, i := range expected {
		if !bst.BstPut(types.Int(i), types.Int(i)) {
			t.Errorf("Element %v should have been added to the tree", i)
		}
	}

	for _, i := range expected {
		if bst.Find(types.Int(i)) == nil {
			t.Errorf("Element %v expected to be in the tree, but was not", i)
		}
	}

	if bst.BstPut(types.Int(4), 44) {
		t.Error("Duplicate elements should not be added")
	}

	if bst.Find(types.Int(4)) == 44 {
		t.Error("Previously inserted elements should not be updated")
	}

	if c := bst.count; c != len(expected) {
		t.Errorf("Tree expected to have %v elements, but has %v instead", len(expected), c)
	}
}

func TestRemove_SingleElement(t *testing.T) {
	bst := NewBST()

	bst.BstPut(types.Int(5), 10)

	if !bst.Delete(types.Int(5)) {
		t.Errorf("Element %v should have been removed", 5)
	}

	if bst.Find(types.Int(5)) != nil {
		t.Errorf("Element %v should not have been found", 5)
	}
}

func TestRemove_RootWithSingleChild(t *testing.T) {
	bst := NewBST()

	bst.BstPut(types.Int(5), 10)
	bst.BstPut(types.Int(4), 8)

	if !bst.Delete(types.Int(5)) {
		t.Errorf("Element %v should have been removed", 5)
	}

	if bst.Find(types.Int(5)) != nil {
		t.Errorf("Element %v should not have been found", 5)
	}

	if bst.Find(types.Int(4)) == nil {
		t.Errorf("Element with key %v was not found", 4)
	}

	if bst.count != 1 {
		t.Errorf("Expected element count %v found cound %v", 1, bst.count)
	}
}

func TestRemove_RootWithTwoChildren(t *testing.T) {
	bst := NewBST()

	bst.BstPut(types.Int(5), 10)
	bst.BstPut(types.Int(4), 8)
	bst.BstPut(types.Int(6), 12)

	if !bst.Delete(types.Int(5)) {
		t.Errorf("Element %v should have been removed", 5)
	}

	if bst.Find(types.Int(5)) != nil {
		t.Errorf("Element %v should not have been found", 5)
	}

	if bst.Find(types.Int(4)) == nil {
		t.Errorf("Element with key %v was not found", 4)
	}

	if bst.Find(types.Int(6)) == nil {
		t.Errorf("Element with key %v was not found", 6)
	}

	if bst.count != 2 {
		t.Errorf("Expected element count %v found cound %v", 1, bst.count)
	}
}

func TestRemove(t *testing.T) {
	expected := []int{5, 3, 7, 4, 6}
	bst := NewBST()

	for _, i := range expected {
		bst.BstPut(types.Int(i), i)
	}

	if !bst.Delete(types.Int(6)) {
		t.Errorf("Element %v should have been removed from the tree", 6)
	}

	for _, i := range expected[0:3] {
		if bst.Find(types.Int(i)) == nil {
			t.Errorf("Element %v expected to be in the tree, but was not", i)
		}
	}

	if d := expected[len(expected)-1]; bst.Find(types.Int(d)) != nil {
		t.Errorf("Element %v should have been removed from the tree", d)
	}

	if bst.Delete(types.Int(6)) {
		t.Error("Duplicate elements should not be delete")
	}

	if c := bst.count; c != len(expected)-1 {
		t.Errorf("Tree expected to have %v elements, but has %v instead", len(expected)-1, c)
	}
}

func TestTraverse_InOrder(t *testing.T) {
	elements := []int{5, 3, 7, 4, 6}
	expected := []int{3, 4, 5, 6, 7}

	bst := NewBST()

	for _, i := range elements {
		bst.BstPut(types.Int(i), i)
	}

	i := 0
	for e := range bst.Traverse(InOrder) {
		if e != expected[i] {
			t.Errorf("Expected to traverse %v, but instead traversed %v", expected[i], e)
		}
		i++
	}
}

func TestTraverse_PreOrder(t *testing.T) {
	elements := []int{5, 3, 7, 4, 6}
	expected := []int{5, 3, 4, 7, 6}

	bst := NewBST()

	for _, i := range elements {
		bst.BstPut(types.Int(i), i)
	}

	i := 0
	for e := range bst.Traverse(PreOrder) {
		if e != expected[i] {
			t.Errorf("Expected to traverse %v, but instead traversed %v", expected[i], e)
		}
		i++
	}
}

func TestTraverse_PostOrder(t *testing.T) {
	elements := []int{5, 3, 7, 4, 6}
	expected := []int{4, 3, 6, 7, 5}

	bst := NewBST()

	for _, i := range elements {
		bst.BstPut(types.Int(i), i)
	}

	i := 0
	for e := range bst.Traverse(PostOrder) {
		if e != expected[i] {
			t.Errorf("Expected to traverse %v, but instead traversed %v", expected[i], e)
		}
		i++
	}
}

func TestClear(t *testing.T) {
	elements := []int{5, 3, 7, 4, 6}

	bst := NewBST()

	for _, i := range elements {
		bst.BstPut(types.Int(i), i)
	}

	bst.Clear()

	if c := bst.count; c != 0 {
		t.Errorf("Expected tree to be empty, but has %v elements", c)
	}

	if p := bst.String(); len(p) != 0 {
		t.Errorf("No elements expected in the tree, but found %v", p)
	}
}

func (t *BinaryTree) String() (s string) {
	print(t.root, &s)
	return
}

func print(n *BinaryNode, s *string) {
	if n == nil {
		return
	}

	*s += fmt.Sprintf("%p %v\n", n, n)
	print(n.left, s)
	print(n.right, s)
}

func BenchmarkInsert(b *testing.B) {
	bst := NewBST()
	for _, i := range rand.Perm(b.N) {
		bst.BstPut(types.Int(i), i)
	}
}

func BenchmarkDelete(b *testing.B) {
	bst := NewBST()
	for _, i := range rand.Perm(b.N) {
		bst.BstPut(types.Int(i), i)
	}

	b.ResetTimer()
	for _, i := range rand.Perm(b.N) {
		bst.Delete(types.Int(i))
	}
}

func BenchmarkFind(b *testing.B) {
	bst := NewBST()
	for _, i := range rand.Perm(b.N) {
		bst.BstPut(types.Int(i), i)
	}

	b.ResetTimer()
	for _, i := range rand.Perm(b.N) {
		bst.Find(types.Int(i))
	}
}

/**      AVL Related code **/

var rand *mrand.Rand

func init() {
	seed := make([]byte, 8)
	if _, err := crand.Read(seed); err == nil {
		rand = utils.ThreadSafeRand(int64(eb.BigEndian.Uint64(seed)))
	} else {
		panic(err)
	}
}

func randstr(length int) types.String {
	return types.String(utils.RandStr(length))
}

func TestAvlPutHasGetRemove(t *testing.T) {

	type record struct {
		key   types.String
		value types.String
	}

	records := make([]*record, 400)
	var tree *BinaryNode
	var err error
	var val interface{}
	var updated bool

	ranrec := func() *record {
		return &record{randstr(20), randstr(20)}
	}

	for i := range records {
		r := ranrec()
		records[i] = r
		tree, updated = tree.AvlPut(r.key, types.String(""))
		if updated {
			t.Error("should have not been updated")
		}
		tree, updated = tree.AvlPut(r.key, r.value)
		if !updated {
			t.Error("should have been updated")
		}
		if tree.Size() != (i + 1) {
			t.Error("size was wrong", tree.Size(), i+1)
		}
	}

	for _, r := range records {
		if has := tree.Has(r.key); !has {
			t.Error("Missing key")
		}
		if has := tree.Has(randstr(12)); has {
			t.Error("Table has extra key")
		}
		if val, err := tree.Get(r.key); err != nil {
			t.Error(err, val.(types.String), r.value)
		} else if !(val.(types.String)).Equals(r.value) {
			t.Error("wrong value")
		}
	}

	for i, x := range records {
		if tree, val, err = tree.AvlRemove(x.key); err != nil {
			t.Error(err)
		} else if !(val.(types.String)).Equals(x.value) {
			t.Error("wrong value")
		}
		for _, r := range records[i+1:] {
			if has := tree.Has(r.key); !has {
				t.Error("Missing key")
			}
			if has := tree.Has(randstr(12)); has {
				t.Error("Table has extra key")
			}
			if val, err := tree.Get(r.key); err != nil {
				t.Error(err)
			} else if !(val.(types.String)).Equals(r.value) {
				t.Error("wrong value")
			}
		}
		if tree.Size() != (len(records) - (i + 1)) {
			t.Error("size was wrong", tree.Size(), (len(records) - (i + 1)))
		}
	}
}

func TestImmutableAvlPutHasGetRemove(t *testing.T) {

	type record struct {
		key   types.String
		value types.String
	}

	records := make([]*record, 400)
	var tree *ImmutableAvlNode
	var err error
	var val interface{}
	var updated bool

	ranrec := func() *record {
		return &record{randstr(20), randstr(20)}
	}

	for i := range records {
		r := ranrec()
		records[i] = r
		tree, updated = tree.Put(r.key, types.String(""))
		if updated {
			t.Error("should have not been updated")
		}
		tree, updated = tree.Put(r.key, r.value)
		if !updated {
			t.Error("should have been updated")
		}
		if tree.Size() != (i + 1) {
			t.Error("size was wrong", tree.Size(), i+1)
		}
	}

	for _, r := range records {
		if has := tree.Has(r.key); !has {
			t.Error("Missing key")
		}
		if has := tree.Has(randstr(12)); has {
			t.Error("Table has extra key")
		}
		if val, err := tree.Get(r.key); err != nil {
			t.Error(err, val.(types.String), r.value)
		} else if !(val.(types.String)).Equals(r.value) {
			t.Error("wrong value")
		}
	}

	for i, x := range records {
		if tree, val, err = tree.Remove(x.key); err != nil {
			t.Error(err)
		} else if !(val.(types.String)).Equals(x.value) {
			t.Error("wrong value")
		}
		for _, r := range records[i+1:] {
			if has := tree.Has(r.key); !has {
				t.Error("Missing key")
			}
			if has := tree.Has(randstr(12)); has {
				t.Error("Table has extra key")
			}
			if val, err := tree.Get(r.key); err != nil {
				t.Error(err)
			} else if !(val.(types.String)).Equals(r.value) {
				t.Error("wrong value")
			}
		}
		if tree.Size() != (len(records) - (i + 1)) {
			t.Error("size was wrong", tree.Size(), (len(records) - (i + 1)))
		}
	}
}

func TestIterators(t *testing.T) {
	var data []int = []int{
		1, 5, 7, 9, 12, 13, 17, 18, 19, 20,
	}
	var order []int = []int{
		6, 1, 8, 2, 4, 9, 5, 7, 0, 3,
	}

	test := func(tree types.TreeMap) {
		t.Logf("%T", tree)
		for j := range order {
			if err := tree.Put(types.Int(data[order[j]]), order[j]); err != nil {
				t.Error(err)
			}
		}

		j := 0
		for k, v, next := tree.Iterate()(); next != nil; k, v, next = next() {
			if !k.Equals(types.Int(data[j])) {
				t.Error("Wrong key")
			}
			if v.(int) != j {
				t.Error("Wrong value")
			}
			j += 1
		}

		j = 0
		for k, next := tree.Keys()(); next != nil; k, next = next() {
			if !k.Equals(types.Int(data[j])) {
				t.Error("Wrong key")
			}
			j += 1
		}

		j = 0
		for v, next := tree.Values()(); next != nil; v, next = next() {
			if v.(int) != j {
				t.Error("Wrong value")
			}
			j += 1
		}
	}
	test(NewBinaryTree(AVL))
	test(NewImmutableAvlTree())
}

func BenchmarkAvlTree(b *testing.B) {
	b.StopTimer()

	type record struct {
		key   types.String
		value types.String
	}

	records := make([]*record, 100)

	ranrec := func() *record {
		return &record{randstr(20), randstr(20)}
	}

	for i := range records {
		records[i] = ranrec()
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		t := NewBinaryTree(AVL)
		for _, r := range records {
			t.Put(r.key, r.value)
		}
		for _, r := range records {
			t.Remove(r.key)
		}
	}
}

func BenchmarkImmutableAvlTree(b *testing.B) {
	b.StopTimer()

	type record struct {
		key   types.String
		value types.String
	}

	records := make([]*record, 100)

	ranrec := func() *record {
		return &record{randstr(20), randstr(20)}
	}

	for i := range records {
		records[i] = ranrec()
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		t := NewImmutableAvlTree()
		for _, r := range records {
			t.Put(r.key, r.value)
		}
		for _, r := range records {
			t.Remove(r.key)
		}
	}
}
