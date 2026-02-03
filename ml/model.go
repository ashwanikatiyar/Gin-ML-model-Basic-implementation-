package ml

import "math"


type LogisticRegression struct{
	Weights []float64
	Bias float64
}

func NewLogisticRegression(size int)*LogisticRegression{
	return &LogisticRegression{
		Weights: make([]float64 , size),
		Bias: 0.0,
	}
}

func sigmoid(x float64) float64{
	return 1.0 / (1.0 + math.Exp(-x))
}

func(lr *LogisticRegression) PredictProb(x []float64) float64{

	sum := lr.Bias
	for i ,val := range x{
		sum += lr.Weights[i]*val
	}
	return sigmoid(sum)
}