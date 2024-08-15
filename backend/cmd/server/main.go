package main

import (
	"log"
    "github.com/srujanr40/cal/backend/internal/db"
    "github.com/srujanr40/cal/backend/internal/handlers"
    "github.com/srujanr40/cal/backend/pkg/config"

	"net/http"
	"github.com/gin-gonic/gin"
)

// type health struct {
// 	ID     string `json:"id"`
// 	Name   string `json:"name"`
// 	Status string `json:"status"`
// }

// var healths = []health{
// 	{
// 		ID:     "1",
// 		Name:   "test",
// 		Status: "ok",
// 	},
// }

// func getHealths(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, healths)
// }

func main() {
	// Set up Gin router
    router := gin.Default()

	// health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "I'm alive!")
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// router.GET("/healths", getHealths)

	// Load configuration (e.g., AWS region, credentials)
    cfg := config.LoadConfig()

    // Initialize the DynamoDB client
    dbClient := db.NewDynamoClient(cfg)

    // Define routes
    router.GET("/items", handlers.GetItemsHandler(dbClient))
    router.POST("/items", handlers.CreateItemHandler(dbClient))

    // Start the server
    log.Println("Server is running on port 8080")
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}