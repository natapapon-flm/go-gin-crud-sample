package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/natapapon-flm/go-gin/controllers"
)

func SetupItemRoutes(routers *gin.Engine) {
	v1 := routers.Group("api/v1");
	items := v1.Group("items")
	items.GET("", controllers.GetAllItems)
	items.GET("/:Id", controllers.GetItemsById)
	items.POST("", controllers.Create)
	items.PATCH("/:Id", controllers.Update)
	items.DELETE("/:Id", controllers.Delete)
}