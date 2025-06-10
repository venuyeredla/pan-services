package types

type Equatable interface {
	Equals(b Equatable) bool
}

type Sortable interface {
	Equatable
	Less(b Sortable) bool
}

type Hashable interface {
	Sortable
	Hash() int
}

type Iterator func() (item interface{}, next Iterator)

type Iterable interface {
	Iterate() Iterator
}

type Marshaler interface {
	MarshalBinary() (data []byte, err error)
	UnmarshalBinary(data []byte) error
}

type MHashable interface {
	Hashable
	Marshaler
}

type ItemMarshal func(Hashable) ([]byte, error)
type ItemUnmarshal func([]byte) (Hashable, error)

type KIterator func() (key Hashable, next KIterator)
type KVIterator func() (key Hashable, value interface{}, next KVIterator)
type Coroutine func(send interface{}) (recv interface{}, next Coroutine)

type KIterable interface {
	Keys() KIterator
}

type VIterable interface {
	Values() Iterator
}

type KVIterable interface {
	Iterate() KVIterator
}

type MapIterable interface {
	KIterable
	VIterable
	KVIterable
}

type Sized interface {
	Size() int
}

type MapOperable interface {
	Sized
	Has(key Hashable) bool
	Put(key Hashable, value interface{}) (err error)
	Get(key Hashable) (value interface{}, err error)
	Remove(key Hashable) (value interface{}, err error)
}

type WhereFunc func(value interface{}) bool

type MultiMapOperable interface {
	Sized
	Has(key Hashable) bool
	Count(key Hashable) int
	Add(key Hashable, value interface{}) (err error)
	Replace(key Hashable, where WhereFunc, value interface{}) (err error)
	Find(key Hashable) KVIterator
	RemoveWhere(key Hashable, where WhereFunc) (err error)
}

type Map interface {
	MapIterable
	MapOperable
}

type MultiMap interface {
	MapIterable
	MultiMapOperable
}

type MultiMapTemp interface {
	Keys() (ItemIterator, error)
	Values() (ItemIterator, error)
	Iterate() (Iterator, error)
	Backward() (Iterator, error)
	Find(key []byte) (Iterator, error)
	DoFind(key []byte, do func([]byte, []byte) error) error
	Range(from, to []byte) (Iterator, error)
	DoRange(from, to []byte, do func([]byte, []byte) error) error
	Has(key []byte) (bool, error)
	Count(key []byte) (int, error)
	Add(key []byte, value []byte) error
	Remove(key []byte, where func([]byte) bool) error
	Size() int
}

type ContainerOperable interface {
	Has(item Hashable) bool
}

type ItemsOperable interface {
	Sized
	ContainerOperable
	Item(item Hashable) (Hashable, error)
	Add(item Hashable) (err error)
	Delete(item Hashable) (err error)
	Extend(items KIterator) (err error)
}

type OrderedOperable interface {
	Get(i int) (item Hashable, err error)
	Find(item Hashable) (idx int, has bool, err error)
}

type ListIterable interface {
	Items() KIterator
}

type IterableContainer interface {
	Sized
	ListIterable
	Has(item Hashable) bool
}

type StackOperable interface {
	Push(item Hashable) (err error)
	Pop() (item Hashable, err error)
}

type DequeOperable interface {
	EnqueFront(item Hashable) (err error)
	EnqueBack(item Hashable) (err error)
	DequeFront() (item Hashable, err error)
	DequeBack() (item Hashable, err error)
	First() (item Hashable)
	Last() (item Hashable)
}

type LinkedOperable interface {
	Sized
	ContainerOperable
	StackOperable
	DequeOperable
}

type ListOperable interface {
	Sized
	ContainerOperable
	Append(item Hashable) (err error)
	Get(i int) (item Hashable, err error)
	Set(i int, item Hashable) (err error)
	Insert(i int, item Hashable) (err error)
	Remove(i int) (err error)
}

type OrderedList interface {
	ListIterable
	OrderedOperable
}

type List interface {
	ListIterable
	ListOperable
}

type HList interface {
	Hashable
	List
}

type Stack interface {
	Sized
	ContainerOperable
	StackOperable
}

type Deque interface {
	Sized
	ContainerOperable
	DequeOperable
}

type Linked interface {
	LinkedOperable
	ListIterable
}

type Tree interface {
	Root() TreeNode
}

type TreeMap interface {
	Tree
	Map
}

type TreeNode interface {
	Key() Hashable
	Value() interface{}
	Children() TreeNodeIterator
	GetChild(int) TreeNode // if your tree can't support this simply panic
	// many of the utility functions do not require this
	// however, it is recommended that you implement it
	// if possible (for instance, post-order traversal
	// requires it).
	ChildCount() int // a negative value indicates this tree can't provide
	// an accurate count.
}
type TreeNodeIterator func() (node TreeNode, next TreeNodeIterator)

type BinaryTreeNode interface {
	TreeNode
	Left() BinaryTreeNode
	Right() BinaryTreeNode
}

type ItemIterator func() ([]byte, error, ItemIterator)

/*
	func Do(run func() (Iterator, error), do func(key []byte, value []byte) error) error {
		kvi, err := run()
		if err != nil {
			return err
		}
		var key []byte
		var value []byte
		for key, value, err, kvi = kvi(); kvi != nil; key, value, err, kvi = kvi() {
			e := do(key, value)
			if e != nil {
				return e
			}
		}
		return err
	}
*/
func DoItem(run func() (ItemIterator, error), do func([]byte) error) error {
	it, err := run()
	if err != nil {
		return err
	}
	var item []byte
	for item, err, it = it(); it != nil; item, err, it = it() {
		e := do(item)
		if e != nil {
			return e
		}
	}
	return err
}
