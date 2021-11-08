package routers

import (
	"store-management-api/controllers"

	"github.com/gin-gonic/gin"
)

func StoreRouter(app *gin.Engine) {
	app.GET("/api/inventory", controllers.GetInventories())
	app.GET("/api/inventory/get/:inventory_id", controllers.GetInventory())
	app.POST("/api/inventory/create", controllers.CreateInventory())
}
