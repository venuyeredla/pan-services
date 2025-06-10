package types

import (
	"bytes"
	"hash/fnv"
)

type String string
type ByteSlice []byte

func (str String) Equals(other Equatable) bool {
	if o, ok := other.(String); ok {
		return str == o
	} else {
		return false
	}
}

func (str String) Less(other Sortable) bool {
	if o, ok := other.(String); ok {
		return str < o
	} else {
		return false
	}
}

func (str String) Hash() int {
	h := fnv.New32a()
	h.Write([]byte(string(str)))
	return int(h.Sum32())
}

func (str ByteSlice) Equals(other Equatable) bool {
	if o, ok := other.(ByteSlice); ok {
		return bytes.Equal(str, o)
	} else {
		return false
	}
}

func (str ByteSlice) Less(other Sortable) bool {
	if o, ok := other.(ByteSlice); ok {
		return bytes.Compare(str, o) < 0 // -1 if a < b
	} else {
		return false
	}
}

func (str ByteSlice) Hash() int {
	h := fnv.New32a()
	h.Write([]byte(str))
	return int(h.Sum32())
}

func (str *String) MarshalBinary() ([]byte, error) {
	return []byte(*str), nil
}

func (str *String) UnmarshalBinary(data []byte) error {
	*str = String(data)
	return nil
}

func (str *ByteSlice) MarshalBinary() ([]byte, error) {
	return []byte(*str), nil
}

func (str *ByteSlice) UnmarshalBinary(data []byte) error {
	*str = ByteSlice(data)
	return nil
}
