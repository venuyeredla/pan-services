package ai

import (
	"fmt"
	"math"
)

func NuralNetwork() {
	A := NewTensor(3, 3, 2.0)
	A.SetXY(0, 0, 0.9)
	A.SetXY(0, 1, 0.3)
	A.SetXY(0, 2, 0.4)
	A.SetXY(1, 0, 0.2)
	A.SetXY(1, 1, 0.8)
	A.SetXY(1, 2, 0.2)
	A.SetXY(2, 0, 0.1)
	A.SetXY(2, 1, 0.5)
	A.SetXY(2, 2, 0.6)

	b := NewVector(3, 2.0)
	b.SetVectorX(0, 0.9)
	b.SetVectorX(1, 0.1)
	b.SetVectorX(2, 0.8)
	A.Print()
	b.Print()

	c, e := A.Transform(b)
	if e == nil {
		c.Print()
	}
	for i, val := range c.Vector() {
		y := SigMoid(val)
		c.SetVectorX(i, float32(y))
	}

	fmt.Println("Activatin neurons")
	c.Print()

	H := NewTensor(3, 3, 2.0)
	H.SetXY(0, 0, 0.3)
	H.SetXY(0, 1, 0.7)
	H.SetXY(0, 2, 0.5)
	H.SetXY(1, 0, 0.6)
	H.SetXY(1, 1, 0.5)
	H.SetXY(1, 2, 0.2)
	H.SetXY(2, 0, 0.8)
	H.SetXY(2, 1, 0.1)
	H.SetXY(2, 2, 0.9)

	c, e = H.Transform(c)
	if e == nil {
		c.Print()
	}
	for i, val := range c.Vector() {
		y := SigMoid(val)
		c.SetVectorX(i, float32(y))
	}

	fmt.Println("Activatin neurons")
	c.Print()

}

func SigMoid(x float32) float64 {
	y := 1 / (1 + 1/(math.Exp(float64(x))))
	return y
}
