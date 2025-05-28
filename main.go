package main

import (
	"bakeryapp/config"
	"bakeryapp/logger"
	"bakeryapp/routes"
	"strconv"

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
	config.LoadConfig()

	logger.InitLogger()

	logger.Log.Info("Запуск приложения")
	logger.Log.Infof("ENV: %s, PORT: %d", config.Cfg.App.Env, config.Cfg.App.Port)

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
	routes.RegisterDictRoutes(r, db)
	routes.RegisterDocumentRoutes(r, db)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := strconv.Itoa(config.Cfg.App.Port)
	logger.Log.Info("Starting server on port:", port)
	r.Run(":" + port)
}
