package main

import (
	"github.com/gin-gonic/gin"
	"github.com/natapapon-flm/go-gin/database"
	"github.com/natapapon-flm/go-gin/routes"
)

func main() {
	routers := gin.Default()
	

	// url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	// routers.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	routes.SetupItemRoutes(routers)
	database.ConnectDatabase()


	routers.Run()
}