package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Inventory struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `json:"name"`
	Address     string             `json:"address"`
	Email       string             `json:"email"`
	Weight      string             `json:"weight"`
	CreatedAt   time.Time          `json:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt"`
	InventoryId string             `json:"inventoryId"`
}
