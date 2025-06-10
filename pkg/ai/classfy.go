package ai

import "fmt"

type Classifier interface {
	train()
	predict()
}

type LinearModel struct {
	weights *Vector
}

func (lm *LinearModel) train(inputs [][]int, targets []int) {
	if len(inputs) > 0 {
		firstData := inputs[0]
		length := len(firstData)
		weights := NewVector(length, 0.5)
		weights.Set(1, -0.5)
		aggregate := NewVector(length, 0.0)
		for i, value := range inputs {
			iVector := getVector(value)
			targeted := targets[i]
			predicted := VectorDotProduct(weights, iVector)
			updateWeight := false
			if targeted < 0 && predicted > 0 {
				updateWeight = true
			} else if predicted < 0 {
				updateWeight = true
			}
			fmt.Printf("Pre Weights : %v ==> ", weights.Data)
			if updateWeight {
				deviation := -predicted
				fraction := deviation / float32(length)
				for i, weight := range weights.Data {
					wnew := (fraction * weight) / (weight * iVector.Get(i))
					aggregate.Set(i, wnew)
				}
			} else {
				for i, weight := range weights.Data {
					aggregate.Set(i, weight)
				}
			}
			fmt.Printf(" Post Weights : %v\n", weights.Data)
		}

		for i, w := range aggregate.Data {
			avg := w / float32(6)
			weights.Set(i, avg)
		}

		fmt.Printf("Weights A : %v \n", weights.Data)
		lm.weights = weights

	}
}

func (lm *LinearModel) pridict(input []int) float32 {
	iVector := getVector(input)
	predict := VectorDotProduct(lm.weights, iVector)
	return predict
}

func getVector(data []int) *Vector {
	iVector := NewVector(len(data))
	for i, val := range data {
		iVector.Data[i] = float32(val)
	}
	return iVector
}
