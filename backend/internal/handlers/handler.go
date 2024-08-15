package handlers

import (
    "github.com/srujanr40/cal/backend/internal/db"
    "github.com/srujanr40/cal/backend/internal/models"
    "net/http"

    "github.com/aws/aws-sdk-go-v2/service/dynamodb"
    "github.com/gin-gonic/gin"
)

func GetItemsHandler(client *dynamodb.Client) gin.HandlerFunc {
    return func(c *gin.Context) {
        items, err := db.GetItems(client, "YourTableName")
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, items)
    }
}

func CreateItemHandler(client *dynamodb.Client) gin.HandlerFunc {
    return func(c *gin.Context) {
        var item models.Item
        if err := c.ShouldBindJSON(&item); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        if err := db.CreateItem(client, "YourTableName", item); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"status": "item created successfully"})
    }
}