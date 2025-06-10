package compress

import (
	"fmt"
)

type BitStream struct {
	bytes   []byte
	b       byte
	wbits   int   // holds written bits count
	rembits uint8 // holds remaining no of bits to read
	pointer int   //ponter to read from bytes slice
}

func (stream *BitStream) writeByte(b byte) {
	stream.bytes = append(stream.bytes, b)
}

func (stream *BitStream) readByte() byte {
	if stream.pointer >= len(stream.bytes) {
		fmt.Print("\nEnd of bit stream")
		return 0
	}
	stream.pointer += 1
	return stream.bytes[stream.pointer-1] //Reading first byte
}

func (stream *BitStream) WriteUint16(num uint16) {
	stream.writeByte(byte(num >> 8 & 0xff))
	stream.writeByte(byte(num & 0xff))
}
func (stream *BitStream) ReadUint16() uint16 {
	b1 := stream.readByte() << 8
	b2 := stream.readByte()
	var newNum uint16 = uint16(b1) | uint16(b2)
	return newNum
}

func (stream *BitStream) WriteBit(b byte) {
	bit := b & 1
	//fmt.Print("", bit)
	stream.b = stream.b<<1 | bit
	stream.wbits = stream.wbits + 1
	if stream.wbits == 8 {
		stream.bytes = append(stream.bytes, stream.b)
		stream.b = 0
		stream.wbits = 0
	}
}

func (stream *BitStream) ReadBit() (bool, byte) {
	if stream.rembits == 0 {
		if stream.pointer >= len(stream.bytes) {
			fmt.Print("\nEnd of bit stream")
			return false, 0
		}
		//fmt.Println("Reading new byte index: ", stream.pointer)
		stream.b = stream.bytes[stream.pointer] //Reading first byte
		stream.pointer += 1
		stream.rembits = 8
	}
	val := (uint8(stream.b) >> (stream.rembits - 1)) & 1
	stream.rembits -= 1
	fmt.Println("", val)
	return true, val
}
func (stream *BitStream) close() {
	toBeFilled := 8 - stream.wbits
	if toBeFilled < 8 {
		//fmt.Println("Need to be filld :", toBeFilled)
		for i := 0; i < toBeFilled; i++ {
			stream.WriteBit(0)
		}
	}
}

func (stream *BitStream) size() int {
	return len(stream.bytes)
}

func (stream *BitStream) getBytes() []byte {
	return stream.bytes[0:]
}

func (stream *BitStream) ReadReset() {
	stream.pointer = 0
}

func BitReadWriteTest() {
	stream1 := &BitStream{}
	fmt.Println("Iterating from 1 to 32 and writing bits")
	for i := 1; i <= 47; i++ {
		stream1.WriteBit(byte(i))
	}
	stream1.close()
	fmt.Println("")
	stream1.pointer = 0
	stream1.rembits = 0
	p, val := stream1.ReadBit()
	for p {
		fmt.Print("", val)
		p, val = stream1.ReadBit()
	}
}
