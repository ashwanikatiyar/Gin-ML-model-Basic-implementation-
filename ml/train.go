package ml



func(lr *LogisticRegression) Train(
	X [][]float64,
	y []float64,
	epochs int,
	lrRate float64,
){
	for epoch := 0; epoch <epochs ; epoch++{
		for i , x := range X{
			pred := lr.PredictProb(x)
			err := pred - y[i]

			//Update Weights
			for j := range lr.Weights{
				lr.Weights[j] -= lrRate* err * x[j]
			}

			//Update Bias
			lr.Bias -= lrRate * err
		}
	}
}