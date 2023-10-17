package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/natapapon-flm/go-gin/database"
	"github.com/natapapon-flm/go-gin/routes"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	routers := gin.Default()
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error Loading ENV %v", err)
	}
	// url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	// routers.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	routers.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url));
	
	routes.SetupItemRoutes(routers)
	database.ConnectDatabase()


	routers.Run()
}