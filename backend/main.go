package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type health struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

var healths = []health{
	{
		ID:     "1",
		Name:   "test",
		Status: "ok",
	},
}

func getHealths(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, healths)
}

func main() {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "I'm alive!")
	})

	router.GET("/healths", getHealths)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	router.Run("localhost:8080")
}