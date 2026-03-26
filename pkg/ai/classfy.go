package ai

import "fmt"

type Classifier interface {
	train()
	predict()
}

type LinearModel struct {
	Model *Tensor // Weight vector
}

func (lm *LinearModel) train(inputs [][]int, targets []int) {
	if len(inputs) > 0 {
		firstData := inputs[0]
		length := len(firstData)
		weights := NewTensor(1, length, 0.5)
		weights.Vector()[1] = -0.5
		aggregate := NewTensor(1, length, 0.0)
		for i, value := range inputs {
			iVector := getVector(value)
			targeted := targets[i]
			predicted := weights.VectorDotProduct(iVector)
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
				for i, weight := range weights.Vector() {
					wnew := (fraction * weight) / (weight * iVector.Vector()[i])
					aggregate.Vector()[i] = wnew

				}
			} else {
				for i, weight := range weights.Vector() {
					aggregate.Vector()[i] = weight
				}
			}
			fmt.Printf(" Post Weights : %v\n", weights.Data)
		}

		for i, w := range aggregate.Vector() {
			avg := w / float32(6)
			weights.Vector()[i] = avg
		}

		fmt.Printf("Weights A : %v \n", weights.Data)
		lm.Model = weights

	}
}

func (lm *LinearModel) Predict(input []int) float32 {
	iVector := getVector(input)
	predict := lm.Model.VectorDotProduct(iVector)
	return predict
}

func getVector(data []int) *Tensor {
	iVector := NewVector(len(data))
	for i, val := range data {
		iVector.Vector()[i] = float32(val)
	}
	return iVector
}
