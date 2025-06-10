package ai

import (
	"fmt"
	"math"
)

func NuralNetwork() {
	A := NewMatrix(3, 3, 2.0)
	A.Set(0, 0, 0.9)
	A.Set(0, 1, 0.3)
	A.Set(0, 2, 0.4)
	A.Set(1, 0, 0.2)
	A.Set(1, 1, 0.8)
	A.Set(1, 2, 0.2)
	A.Set(2, 0, 0.1)
	A.Set(2, 1, 0.5)
	A.Set(2, 2, 0.6)

	b := NewVector(3, 2.0)
	b.Set(0, 0.9)
	b.Set(1, 0.1)
	b.Set(2, 0.8)
	A.Print()
	b.Print()

	c, e := Transform(A, b)
	if e == nil {
		c.Print()
	}
	for i, val := range c.Data {
		y := SigMoid(val)
		c.Set(i, float32(y))
	}

	fmt.Println("Activatin neurons")
	c.Print()

	H := NewMatrix(3, 3, 2.0)
	H.Set(0, 0, 0.3)
	H.Set(0, 1, 0.7)
	H.Set(0, 2, 0.5)
	H.Set(1, 0, 0.6)
	H.Set(1, 1, 0.5)
	H.Set(1, 2, 0.2)
	H.Set(2, 0, 0.8)
	H.Set(2, 1, 0.1)
	H.Set(2, 2, 0.9)

	c, e = Transform(H, c)
	if e == nil {
		c.Print()
	}
	for i, val := range c.Data {
		y := SigMoid(val)
		c.Set(i, float32(y))
	}

	fmt.Println("Activatin neurons")
	c.Print()

}

func SigMoid(x float32) float64 {
	y := 1 / (1 + 1/(math.Exp(float64(x))))
	return y
}
