package ai

import (
	"errors"
	"fmt"
	"math"
)

type Shape struct {
	Rows int
	Cols int
}

type TensorInterface interface {
	Shape()
	Print()
	Vector()
}

// Tensor of size 3
type Tensor struct {
	Data  [][]float32
	Shape Shape
}

func NewTensor(rows, columns int, defa ...float32) *Tensor {
	var tensor [][]float32
	tensor = make([][]float32, rows)
	for i := 0; i < rows; i++ {
		vector := make([]float32, columns)
		if len(defa) > 0 {
			for j := range vector {
				vector[j] = defa[0]
			}
		}
		tensor[i] = vector
	}
	return &Tensor{Data: tensor, Shape: Shape{Rows: rows, Cols: columns}}
}

func NewVector(dim int, defa ...float32) *Tensor {
	if len(defa) > 0 {
		return NewTensor(1, dim, defa[0])
	}
	return NewTensor(1, dim)
}

func (tesnor *Tensor) Vector() []float32 {
	return tesnor.Vector()
}

func (tesnor *Tensor) SetVectorX(pos int, val float32) {
	tesnor.Vector()[pos] = val
}

func (tesnor *Tensor) SetXY(row int, col int, val float32) {
	tesnor.Data[row][col] = val
}

func (tensor *Tensor) Print() {
	fmt.Printf("Tensor: [\n")
	for _, rowVector := range tensor.Data {
		fmt.Println(rowVector)
		fmt.Printf("\n")
	}
	fmt.Print("]")
}

func (tensor *Tensor) Rows() int {
	return tensor.Shape.Rows
}

func (tensor *Tensor) Columns() int {
	return tensor.Shape.Cols
}

func (tensorA *Tensor) VectorDotProduct(b *Tensor) float32 {
	if tensorA.Columns() == b.Columns() {
		var scalar float32 = 0.0
		for i := 0; i < tensorA.Columns(); i++ {
			scalar = scalar + tensorA.Vector()[i]*b.Vector()[i]
		}
		return scalar
	}
	return 0.0
}

func (matrix *Tensor) Transform(vector *Tensor) (*Tensor, error) {
	if matrix.Columns() == vector.Shape.Cols {
		var c []float32 = make([]float32, matrix.Shape.Rows)
		for i := 0; i < matrix.Rows(); i++ {
			row := matrix.Data[i]
			var scalar float32 = 0.0
			for j, val := range row {
				scalar = scalar + val*vector.Vector()[j]
			}
			c[i] = scalar
		}
		return &Tensor{Data: [][]float32{c}, Shape: Shape{Rows: 1, Cols: matrix.Shape.Rows}}, nil
	} else {
		errorMsg := fmt.Sprintf("Invalid shapes for multiliplcaiton a=%v, b=%v", matrix.Rows(), vector.Columns())
		return &Tensor{}, errors.New(errorMsg)
	}
}

type Mode interface {
	train()
	predict()
}

// d=sqrt(( pow(x2-x1,2) + pow(y2-y1,1,2))
func Distance(a, b Tensor) float32 {
	if len(a.Vector()) == len(b.Vector()) {
		var defsquare float64 = 0.0

		for i := 0; i < a.Shape.Cols; i++ {
			defsquare = math.Pow(float64(a.Vector()[i]-b.Vector()[i]), 2)
		}
		distance := math.Sqrt(defsquare)
		return float32(distance)
	}
	return 0.0
}

func MatrixAdd(a [][]int, b [][]int) {

}

func MatrixMultiplication(a, b *Tensor) (*Tensor, error) {
	if a.Columns() == b.Rows() {
		cRows := a.Rows()
		cCols := b.Columns()
		var c [][]float32 = make([][]float32, cRows)
		for i := 0; i < cRows; i++ {
			c[i] = make([]float32, cCols)
			for j := 0; j < cCols; j++ {
				var cij float32 = 0.0
				for k := 0; k < len(a.Vector()); k++ {
					cij = cij + a.Data[i][k]*b.Data[k][j]
				}
				c[i][j] = cij
			}
		}
		return &Tensor{Data: c, Shape: Shape{Rows: cRows, Cols: cCols}}, nil
	} else {
		errorMsg := fmt.Sprintf("Invalid shapes for multiliplcaiton a=%v, b=%v", a.Rows(), b.Columns())
		return nil, errors.New(errorMsg)

	}

}
