package maths

import (
	"fmt"
	"testing"
)

func Input() [5]int {
	return [5]int{1, 8, 9, 15}
}

func TestBinaryOperators(t *testing.T) {
	var n byte = 10
	var Negation byte = ^n
	// for Unsigned x, ~x =>  255-x
	fmt.Printf("~ %v (unsigned) => ^%08b = %08b => %v\n", n, n, Negation, Negation)
	var n1 int8 = 11
	var Negation2 byte = ^n
	fmt.Printf("~ %v (Signed) => ^%08b = %08b -> %v \n", n1, n1, ^n1, Negation2)
}

func TestConversions(t *testing.T) {
	input := 11
	binString := ToBinary(input)
	decimal := ToDecimal(binString)
	fmt.Printf(" %v => %v  =>  %v \n", input, binString, decimal)

	base := 8
	input = 12345
	binString = ToBase(input, base)
	fmt.Printf(" %v (%v) =>  %v \n", input, base, binString)
}

func TestArthimetic(t *testing.T) {
	//Add(13, 7)
	q, r := Divide(7, 2)
	if q != 3 && r != 1 {
		t.Fail()
	} else {
		fmt.Printf("Q =%v and r=%v", q, r)
	}
}

func TestAdd1(t *testing.T) {
	sum := Add(10, 15)
	difference := substraction(15, 10)
	expected := 25
	if sum != expected {
		t.Errorf("Expected %v , Actual=%v ,diff=%v", expected, sum, difference)
	}
}

func TestLSB(t *testing.T) {
	for _, v := range Input() {
		fmt.Printf(" LSB(%08b)=%08b \n", v, LSB(v))
	}
}
