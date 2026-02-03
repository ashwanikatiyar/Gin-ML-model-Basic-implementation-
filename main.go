package main

import (
	"github.com/gin-gonic/gin"
	"sentiment-api/api"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.POST("/train", api.TrainHandler)
	r.POST("/predict", api.PredictHandler)

	r.Run(":8080")
}
