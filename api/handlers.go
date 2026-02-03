package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sentiment-api/ml"
)

type TrainRequest struct {
	Data []struct {
		Text  string  `json:"text"`
		Label float64 `json:"label"`
	} `json:"data"`
}

type PredictRequest struct {
	Text string `json:"text"`
}


func TrainHandler(c *gin.Context) {
	var req TrainRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	texts := []string{}
	labels := []float64{}

	for _, d := range req.Data {
		texts = append(texts, d.Text)
		labels = append(labels, d.Label)
	}

	vectorizer := ml.NewVectorizer()
	vectorizer.Fit(texts)

	X := [][]float64{}
	for _, t := range texts {
		X = append(X, vectorizer.Transform(t))
	}

	model := ml.NewLogisticRegression(len(vectorizer.Vocab))
	model.Train(X, labels, 1000, 0.1)

	ml.GlobalVectorizer = vectorizer
	ml.GlobalModel = model

	c.JSON(http.StatusOK, gin.H{
		"message": "model trained successfully",
		"vocab":   len(vectorizer.Vocab),
	})
}



func PredictHandler(c *gin.Context) {
	if ml.GlobalModel == nil || ml.GlobalVectorizer == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "model not trained",
		})
		return
	}

	var req PredictRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vec := ml.GlobalVectorizer.Transform(req.Text)
	prob := ml.GlobalModel.PredictProb(vec)

	sentiment := "negative"
	if prob >= 0.5 {
		sentiment = "positive"
	}

	c.JSON(http.StatusOK, gin.H{
		"text":       req.Text,
		"probability": prob,
		"sentiment":  sentiment,
	})
}
