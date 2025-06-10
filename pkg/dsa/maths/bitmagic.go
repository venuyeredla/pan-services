package maths

import (
	"strconv"
)

/*
  Negative numbers are stored as compliment of 2. Negative bit string cancat(1,pow(2,N-1)-k)
 Operatos : &(AND) , |(OR) , ^(XOR), ~(Negation)
 let x => x^x =0
 Let x,y,z numbers   x^y=z  => z^x=y or z^y=x
 Flip(x)= ^x
  Even(x)= x&1==0

 Shifting left/right (Unsigned numbers also)
 Bit Masking.

 LSB to zero. let y=1  x=x&(^(y<<postion)) postiton[0-n] . Identify position by left shift of 1 then
 MSB to Zero

Note :: odd or even number of times counting can be done with xor operator
1's compliment : Flipping bits
2's compliment : Flipping bits and adding 1 gives complment.

/* Standard problem
1. Counting bits which set to 1.
2. Clear the lowest set bit.
3. Parity of binary word.
4. Swap bits.
6. Reverse bits.
7. Number with same weight, Number of one's in binary string.
8. Compute x*y or x/y without using arthmetic operators.

*/

// Change Lowest set bit
// Another approach x & (x-1)
// Extraction of lowest set bit x & ^(x-1)
func LSB(x int) int {
	y := 1
	for x&y == 0 {
		y <<= 1
	}
	return x & (^y)
}
func MSB(x int) int {
	y := 1 << 31
	for x&y == 0 {
		y >>= 1
	}
	return x & (^y)
}

func swap(a int, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}

// parity(x) = 1 if count(1's) odd else 0.
func Parity(x int) int {
	parity := 0
	for x > 0 {
		parity = parity ^ (x & 1)
		x = x >> 1
	}
	return parity
}

// Adding 1 to x => x&y==0 you can do x|y
func Add1(x int) int {
	y := 1
	for x&y != 0 {
		x = x & (^y)
		y <<= 1
	}
	return x | y
}

func Add(x, y int) int {
	z := 0
	c := 0
	position := 0
	for x > 0 || y > 0 {
		xb := x & 1
		yb := y & 1

		if c&xb == 1 {
			c <<= 1
		} else {
			c = c ^ xb
		}
		if c&yb == 1 {
			c <<= 1
		} else {
			c = c ^ yb
		}

		z = z | ((c & 1) << position)
		position++
		c >>= 1
		x >>= 1
		y >>= 1

	}
	z = z | ((c & 1) << position)
	return z
}

func substraction(x, y int) int { // Second number 2' complment and adding numbers
	y = ^y
	y += 1
	diff := Add(x, y)
	return MSB(diff)
}

func multiply(x, y int) int {
	sum := 0
	for x > 0 {
		if x&1 == 1 {
			sum = Add(sum, y)
		}
		x >>= 1
		y <<= 1
	}
	return sum
}

func Divide(x, y int) (int, int) {
	q := 0
	for x >= y {
		x -= y
		q++
	}
	return q, x
}

// byte 1000 1111
// Bruteforce = converting as arran and flipping bits
// By using bit masking. i==j bits are no swapping needed.
func SwapBits(x, i, j int) int {
	if ((x >> i) & 1) != ((x >> j) & 1) {
		// ith and jth bits differ need flip bits
		bitMask := (1 << i) | (1 << j)
		x ^= bitMask
	}
	return x
}

func ToBinary(num int) string {

	if num == 0 {
		return "0"
	}
	var bitStr string = ""
	for num > 0 {
		bitStr = strconv.Itoa(num&1) + bitStr //Extracting last  bit
		num = num >> 1
	}
	return bitStr
}

// Binary to Decimal
// TODO : to octal and hex decimal
func ToDecimal(binary string) int {
	num := 0
	for _, val := range binary { // 100 => 4
		if val == '0' {
			num = num<<1 | 0 // Adding binary number to a number
		} else {
			num = num<<1 | 1 // Adding binary number to a number
		}
	}
	return num
}

// Integer to given base.
func ToBase(number int, base int) string {
	bitString := NewBitring()
	for number >= base {
		r := number % base
		number = number / base
		// result = strconv.Itoa(r) + result
		bitString.Append(byte(r))
	}
	bitString.Append(byte(number))
	return bitString.Get()
}

type BitString struct {
	bitstr []byte
}

func NewBitring() *BitString {
	return &BitString{bitstr: make([]byte, 0, 10)}
}

func (bs *BitString) Append(bit byte) {
	bs.bitstr = append(bs.bitstr, bit)
}

func (bs *BitString) Get() string {
	var result string = ""
	for i := len(bs.bitstr) - 1; i >= 0; i-- {
		result = result + strconv.Itoa(int(bs.bitstr[i]))
	}
	return result
}
