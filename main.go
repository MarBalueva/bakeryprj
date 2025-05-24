package main

import (
	"bakeryapp/config"
	"bakeryapp/routes"

	_ "bakeryapp/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Bakery App API
// @version 1.0
// @description REST API для онлайн-магазина кондитерских изделий
// @host localhost
// @schemes https
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	db := config.ConnectDB()
	r := gin.Default()

	routes.RegisterAuthRoutes(r, db)
	routes.RegisterProductRoutes(r, db)
	routes.RegisterUserRoutes(r, db)
	routes.RegisterOrderRoutes(r, db)
	routes.RegisterEmployeeRoutes(r, db)
	routes.RegisterBasketRoutes(r, db)
	routes.RegisterClientRoutes(r, db)
	routes.RegisterPaymentRoutes(r, db)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
