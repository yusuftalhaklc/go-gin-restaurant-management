package routes

import (
	"restaurantmanagement/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderitems", controllers.GetOrderItems())
	incomingRoutes.GET("/orderitems/:orderitem_id", controllers.GetOrderItem())
	incomingRoutes.GET("/orderitems-order/:order_id", controllers.GetOrderItemsByOrder())
	incomingRoutes.POST("/orderitems/", controllers.CreateOrderItem())
	incomingRoutes.PATCH("/orderitems/:orderitem_id", controllers.UpdateOrderItem())
}
