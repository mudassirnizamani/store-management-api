package controllers

import (
	"context"
	"net/http"
	"store-management-api/database"
	"store-management-api/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateInventory() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var inventory models.Inventory

		if err := c.BindJSON(&inventory); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ServerError", "description": err.Error()})
			defer cancel()
			return
		}

		inventory.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		inventory.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		inventory.ID = primitive.NewObjectID()
		inventory.InventoryId = inventory.ID.Hex()

		inventoryInertionNumber, insertError := database.InventoryCollection.InsertOne(ctx, inventory)
		defer cancel()

		if insertError != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ServerError", "description": insertError.Error()})
			return
		}

		c.JSON(http.StatusOK, inventoryInertionNumber)
	}
}

func GetInventories() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		result, err := database.InventoryCollection.Find(context.TODO(), bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ServerError", "description": err.Error()})
			defer cancel()
			return
		}

		var allInventories []bson.M

		if err := result.All(ctx, &allInventories); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ServerError", "description": err.Error()})
			defer cancel()
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, allInventories)
	}
}

func GetInventory() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
