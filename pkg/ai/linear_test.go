package ai

import (
	"testing"
)

func AssertTrue(t *testing.T, actual, expected float32) {
	if actual != expected {
		t.Logf("Expected=%v but Actual =%v ", expected, actual)
		t.FailNow()
	}
}

func TestVector(t *testing.T) {
	//t.Skip()
	cVector := ColVector(5)
	cVector.Print()
	a := NewVector(3, 2.0)
	b := NewVector(3, 2.0)
	a.Print()
	b.Print()
	scalar := VectorDotProduct(a, b)
	AssertTrue(t, scalar, 12)
}

func TestVectorTransformation(t *testing.T) {

	A := NewMatrix(5, 2, 2.0)
	b := NewVector(2, 2.0)
	A.Print()
	b.Print()

	cMatrix, error := Transform(A, b)
	if error == nil {
		cMatrix.Print()
	} else {
		t.Error(error.Error())
	}
}

func TestMatrixMuliplication(t *testing.T) {
	A := NewMatrix(5, 2, 2.0)
	B := NewMatrix(2, 1, 2.0)
	A.Print()
	B.Print()
	cMatrix, error := MatrixMultiplication(A, B)
	if error == nil {
		cMatrix.Print()
	} else {
		t.Error(error.Error())
	}
}
