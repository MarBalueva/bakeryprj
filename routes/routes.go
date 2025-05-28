package routes

import (
	"bakeryapp/controllers"
	"bakeryapp/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterProductRoutes(r *gin.Engine, db *gorm.DB) {
	productGroup := r.Group("/products")
	{
		productGroup.GET("/", controllers.GetProducts(db))
		productGroup.GET("/:id", controllers.GetProductByID(db))
		productGroup.POST("/", middleware.RoleMiddleware(db, "admin", "manager"), controllers.CreateProduct(db))
		productGroup.PUT("/:id", middleware.RoleMiddleware(db, "admin", "manager"), controllers.UpdateProduct(db))
		productGroup.DELETE("/:id", middleware.RoleMiddleware(db, "admin", "manager"), controllers.DeleteProduct(db))
	}
}

func RegisterAuthRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/register", controllers.Register(db))
	r.POST("/login", controllers.Login(db))

	auth := r.Group("/protected")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Authorized access"})
		})
	}
}

func RegisterUserRoutes(router *gin.Engine, db *gorm.DB) {
	users := router.Group("/users")
	{
		users.POST("", middleware.RoleMiddleware(db, "admin"), controllers.CreateUser(db))
		users.PUT("/:id", middleware.RoleMiddleware(db, "admin"), controllers.UpdateUser(db))
		users.DELETE("/:id", middleware.RoleMiddleware(db, "admin"), controllers.DeleteUser(db))
		users.GET("", middleware.RoleMiddleware(db, "admin"), controllers.GetAllUsers(db))
		users.GET("/:id", middleware.RoleMiddleware(db, "admin"), controllers.GetUserByID(db))
		users.GET("/:id/access-groups", middleware.RoleMiddleware(db, "admin"), controllers.GetUserAccessGroups(db))
		users.PUT("/:id/access-groups", middleware.RoleMiddleware(db, "admin"), controllers.UpdateUserAccessGroups(db))
	}
}

func RegisterEmployeeRoutes(router *gin.Engine, db *gorm.DB) {
	employees := router.Group("/employees")
	{
		employees.POST("", middleware.RoleMiddleware(db, "admin"), controllers.CreateEmployee(db))
		employees.PUT("/:id", middleware.RoleMiddleware(db, "admin"), controllers.UpdateEmployee(db))
		employees.PUT("/:id/fire", middleware.RoleMiddleware(db, "admin"), controllers.FireEmployee(db))
		employees.DELETE("/:id", middleware.RoleMiddleware(db, "admin"), controllers.DeleteEmployee(db))
		employees.GET("", middleware.RoleMiddleware(db, "admin", "manager"), controllers.GetAllEmployees(db))
		employees.GET("/:id", middleware.RoleMiddleware(db, "admin", "manager"), controllers.GetEmployeeById(db))
	}
}

func RegisterOrderRoutes(router *gin.Engine, db *gorm.DB) {
	orders := router.Group("/orders")
	{
		orders.POST("", middleware.RoleMiddleware(db, "admin", "manager", "client"), controllers.CreateOrder(db))
		orders.GET("", middleware.RoleMiddleware(db, "admin", "manager", "client"), controllers.GetAllOrders(db))
		orders.GET("/:id", middleware.RoleMiddleware(db, "admin", "manager", "client"), controllers.GetOrderById(db))
		orders.GET("/:id/status", middleware.RoleMiddleware(db, "admin", "manager", "client"), controllers.GetOrderStatus(db))
		orders.PUT("/:id", middleware.RoleMiddleware(db, "admin", "manager", "client"), controllers.UpdateOrder(db))
		orders.PUT("/:id/status", middleware.RoleMiddleware(db, "admin", "manager"), controllers.UpdateOrderStatus(db))
		orders.DELETE("/:id", middleware.RoleMiddleware(db, "admin", "manager", "client"), controllers.DeleteOrder(db))
	}
}

func RegisterClientRoutes(router *gin.Engine, db *gorm.DB) {
	clients := router.Group("/clients")
	{
		clients.POST("", middleware.RoleMiddleware(db, "admin", "manager", "client"), controllers.CreateClient(db))
		clients.PUT("/:id", middleware.RoleMiddleware(db, "admin", "manager", "client"), controllers.UpdateClient(db))
		clients.DELETE("/:id", middleware.RoleMiddleware(db, "admin", "manager"), controllers.DeleteClient(db))
		clients.GET("", middleware.RoleMiddleware(db, "admin", "manager"), controllers.GetAllClients(db))
		clients.GET("/:id", middleware.RoleMiddleware(db, "admin", "manager", "client"), controllers.GetClientByID(db))
	}
}

func RegisterBasketRoutes(router *gin.Engine, db *gorm.DB) {
	basket := router.Group("/basket")
	{
		basket.POST("", middleware.RoleMiddleware(db, "client"), controllers.AddProductToBasket(db))
		basket.DELETE("/:clientId/:productId", middleware.RoleMiddleware(db, "client"), controllers.RemoveProductFromBasket(db))
		basket.GET("/:clientId", middleware.RoleMiddleware(db, "admin", "client"), controllers.GetBasketByClientID(db))
	}
}

func RegisterPaymentRoutes(router *gin.Engine, db *gorm.DB) {
	payments := router.Group("/payments")
	{
		payments.POST("", middleware.RoleMiddleware(db, "client", "manager", "admin"), controllers.CreatePayment(db))
		payments.GET("", middleware.RoleMiddleware(db, "manager", "admin"), controllers.GetAllPayments(db))
		payments.GET("/order/:orderId", middleware.RoleMiddleware(db, "client", "manager", "admin"), controllers.GetPaymentsByOrder(db))
		payments.GET("/:id/status", middleware.RoleMiddleware(db, "client", "manager", "admin"), controllers.GetPaymentStatus(db))
		payments.PUT("/:id/status", middleware.RoleMiddleware(db, "manager", "admin"), controllers.UpdatePaymentStatus(db))
	}
}

func RegisterDocumentRoutes(router *gin.Engine, db *gorm.DB) {
	documents := router.Group("/documents")
	{
		documents.POST("", middleware.RoleMiddleware(db, "admin", "manager"), controllers.CreateDocument(db))
		documents.PUT("/:id", middleware.RoleMiddleware(db, "admin", "manager"), controllers.UpdateDocument(db))
		documents.DELETE("/:id", middleware.RoleMiddleware(db, "admin", "manager"), controllers.DeleteDocument(db))
		documents.GET("", middleware.RoleMiddleware(db, "admin", "manager"), controllers.GetAllDocuments(db))
		documents.GET("/:id", middleware.RoleMiddleware(db, "admin", "manager"), controllers.GetDocumentByID(db))
		documents.PUT("/:id/submit", middleware.RoleMiddleware(db, "admin", "manager"), controllers.SubmitDocument(db))
	}
}

func RegisterDictRoutes(router *gin.Engine, db *gorm.DB) {
	dict := router.Group("/dict")
	{
		dict.GET("/:dict", middleware.RoleMiddleware(db, "admin", "manager"), controllers.GetAllDictItems(db))
		dict.POST("/:dict", middleware.RoleMiddleware(db, "admin", "manager"), controllers.CreateDictItem(db))
		dict.DELETE("/:dict/:id", middleware.RoleMiddleware(db, "admin", "manager"), controllers.DeleteDictItem(db))
	}
}
