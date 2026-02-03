package ml

import "strings"


type Vectorizer struct {
	Vocab map[string]int
}

func NewVectorizer() *Vectorizer{
	return &Vectorizer{
		Vocab: make(map[string]int),
	}
}


//Build vocabulary from training data
func (v *Vectorizer) Fit(texts []string){
	index := 0
	for _, text := range texts{
		words := strings.Fields(strings.ToLower(text))
		for _, word := range words{
			if _, exists := v.Vocab[word]; !exists {
				v.Vocab[word] = index
				index ++
			}
		}
	}
}


func(v *Vectorizer) Transform(text string) []float64{

	vec := make([]float64 , len(v.Vocab))
	words := strings.Fields(strings.ToLower(text))

	for _,word := range words{
		if idx , ok := v.Vocab[word]; ok{
			vec[idx]++
		}
	}
	return vec
}