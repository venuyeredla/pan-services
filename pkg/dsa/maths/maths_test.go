package maths

import (
	"fmt"
	"testing"
)

func TestSummations(t *testing.T) {

}

func TestModular(t *testing.T) {
	x := 2
	y := 5
	p := 13
	mod := ModuloPower(x, y, p)
	fmt.Printf("Power is=%v ", mod)
	input := [2]int{7, 9}
	modulo := 11
	ModAdd(input[:], modulo)
	ModMulti(input[:], modulo)
	modExponent := ModExponent(3, 2003, 99)
	fmt.Printf("3, 2003, 99 = %v \n", modExponent)

}
