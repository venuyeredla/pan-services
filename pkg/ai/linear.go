package ai

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Mode interface {
	train()
	predict()
}

type Vector struct {
	Data []float32
}

type Matrix struct {
	Data [][]float32
}

// Tensor of size 3
type Tensor struct {
	Data [][][]float32
}

/*
length : length of vector
defa : default value
*/
func NewVector(length int, defa ...float32) *Vector {
	vec := make([]float32, length)

	if len(defa) > 0 {
		for i := 0; i < length; i++ {
			vec[i] = defa[0]
		}
	}

	return &Vector{Data: vec}
}

func (vector *Vector) Initialize(defa float32) {
	length := vector.Shape().Cols
	for i := 0; i < length; i++ {
		vector.Data[i] = defa
	}
}

func (vector *Vector) Set(i int, v float32) bool {
	vector.Data[i] = v
	return true
}

func (vector *Vector) Get(i int) float32 {
	return vector.Data[i]
}

type Shape struct {
	Rows int
	Cols int
}

func (vector *Vector) Shape() Shape {
	return Shape{Rows: 1, Cols: len(vector.Data)}
}

func (vector *Vector) Print() {
	fmt.Printf("Vector : %v \n", vector.Shape())
	fmt.Println(vector.Data)
	fmt.Print("\n")
}

// Slices are alway one dimensional
func RowVector(length uint) *Matrix {
	return NewMatrix(1, 5)
}

func ColVector(length uint) *Matrix {
	return NewMatrix(5, 1)
}

func (matrix *Matrix) Shape() Shape {
	return Shape{Rows: len(matrix.Data), Cols: len(matrix.Data[0])}
}

func (matrix *Matrix) Print() {
	fmt.Printf("Matrix : %v \n", matrix.Shape())
	for _, row := range matrix.Data {
		fmt.Println(row)
	}
	fmt.Print("\n")
}

func NewMatrix(rows int, columns int, defa ...float32) *Matrix {
	var matrix [][]float32 = make([][]float32, rows)
	for i := 0; i < rows; i++ {
		vector := make([]float32, columns)
		if len(defa) > 0 {
			for j := range vector {
				vector[j] = defa[0]
			}
		}
		matrix[i] = vector
	}
	return &Matrix{Data: matrix}
}

func (matrix *Matrix) Init(defa float32) {
	shape := matrix.Shape()
	rand.Seed(time.Now().UnixMilli())
	for row := 0; row < shape.Rows; row++ {
		for col := 0; col < shape.Cols; col++ {
			matrix.Data[row][col] = defa
		}
	}
}

func (matrix *Matrix) Set(i, j int, v float32) bool {
	matrix.Data[i][j] = v
	return true
}

func (matrix *Matrix) Get(i, j int) float32 {
	return matrix.Data[i][j]
}

func GenTensor(size, rows, columns int) *Tensor {
	var tensor [][][]float32
	tensor = make([][][]float32, size)
	rand.Seed(time.Now().UnixMilli())
	for i := 0; i < size; i++ {
		var matix = make([][]float32, rows)
		for j := 0; j < rows; j++ {
			vector := make([]float32, columns)
			for k := 0; k < columns; k++ {
				vector[j] = rand.Float32()
			}
			matix[j] = vector
		}
		tensor[i] = matix
	}
	return &Tensor{Data: tensor}
}

func (tensor *Tensor) PrintTensor() {
	fmt.Printf("Tensor: [\n")
	for _, matrix := range tensor.Data {
		for _, row := range matrix {
			fmt.Println(row)
		}
		fmt.Printf("\n")
	}
	fmt.Print("]")
}

// d=sqrt(( pow(x2-x1,2) + pow(y2-y1,1,2))
func Distance(a, b Vector) float32 {
	if a.Shape() == b.Shape() {
		var defsquare float64 = 0.0
		for i := 0; i < a.Shape().Cols; i++ {
			defsquare = math.Pow(float64(a.Data[i]-b.Data[i]), 2)
		}
		distance := math.Sqrt(defsquare)
		return float32(distance)
	}
	return 0.0
}

func MatrixAdd(a [][]int, b [][]int) {

}

func VectorDotProduct(a, b *Vector) float32 {
	if a.Shape() == b.Shape() {
		var scalar float32 = 0.0
		for i := 0; i < a.Shape().Cols; i++ {
			scalar = scalar + a.Data[i]*b.Data[i]
		}
		return scalar
	}
	return 0.0
}

func Transform(a *Matrix, b *Vector) (*Vector, error) {
	if a.Shape().Cols == b.Shape().Cols {
		var c []float32 = make([]float32, a.Shape().Rows)
		for i := 0; i < a.Shape().Rows; i++ {
			row := a.Data[i]
			var scalar float32 = 0.0
			for i, val := range row {
				scalar = scalar + val*b.Data[i]
			}
			c[i] = scalar
		}
		return &Vector{Data: c}, nil
	} else {
		errorMsg := fmt.Sprintf("Invalid shapes for multiliplcaiton a=%v, b=%v", a.Shape(), b.Shape())
		return &Vector{}, errors.New(errorMsg)
	}
}

func MatrixMultiplication(a, b *Matrix) (*Matrix, error) {
	if a.Shape().Cols == b.Shape().Rows {
		cRows := a.Shape().Rows
		cCols := b.Shape().Cols
		var c [][]float32 = make([][]float32, cRows)
		for i := 0; i < cRows; i++ {
			c[i] = make([]float32, cCols)
			for j := 0; j < cCols; j++ {
				var cij float32 = 0.0
				for k := 0; k < len(a.Data[0]); k++ {
					cij = cij + a.Data[i][k]*b.Data[k][j]
				}
				c[i][j] = cij
			}
		}
		return &Matrix{Data: c}, nil
	} else {
		errorMsg := fmt.Sprintf("Invalid shapes for multiliplcaiton a=%v, b=%v", a.Shape(), b.Shape())
		return nil, errors.New(errorMsg)
	}

}
