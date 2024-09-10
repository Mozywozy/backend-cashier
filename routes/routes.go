package routes

import (
    "github.com/gin-gonic/gin"
    "cashier-app/controllers"
    "cashier-app/middleware"
)

func RegisterRoutes(router *gin.Engine) {
    // Auth routes
    router.POST("/register", controllers.Register)
    router.POST("/login", controllers.Login)
    router.POST("/logout", middleware.AuthMiddleware(), controllers.Logout)
    
    // Admin routes
    adminRoutes := router.Group("/admin")
    adminRoutes.Use(middleware.AuthMiddleware(), middleware.CheckRoleMiddleware("admin"))
    {
        adminRoutes.GET("/products", controllers.GetProductsAdmin)
        adminRoutes.GET("/products/:id", controllers.GetProductAdmin)
        adminRoutes.POST("/products", controllers.CreateProductAdmin) 
        adminRoutes.PUT("/products/:id", controllers.UpdateProductAdmin)
        adminRoutes.DELETE("/products/:id", controllers.DeleteProductAdmin)
        adminRoutes.GET("/users", controllers.GetUsers)
        adminRoutes.GET("/users/:id", controllers.GetUser)
        adminRoutes.POST("/users", controllers.CreateUser) 
        adminRoutes.PUT("/users/:id", controllers.UpdateUser)
        adminRoutes.DELETE("/users/:id", controllers.DeleteUser)
    }

    // Gudang routes
    gudangRoutes := router.Group("/petugas_gudang")
    gudangRoutes.Use(middleware.AuthMiddleware(), middleware.CheckRoleMiddleware("gudang"))
    {
        gudangRoutes.GET("/products", controllers.GetProducts)
        gudangRoutes.GET("/products/:id", controllers.GetProduct)
        gudangRoutes.POST("/products", controllers.CreateProduct) 
        gudangRoutes.PUT("/products/:id", controllers.UpdateProduct)
        gudangRoutes.DELETE("/products/:id", controllers.DeleteProduct)
    }

    // Kasir routes
    kasirRoutes := router.Group("/petugas_kasir")
    kasirRoutes.Use(middleware.AuthMiddleware(), middleware.CheckRoleMiddleware("kasir"))
    {
        kasirRoutes.GET("/sales", controllers.GetSales)
        kasirRoutes.GET("/sales/:id", controllers.GetSale)
        kasirRoutes.POST("/sales", controllers.CreateSale) 
        kasirRoutes.PUT("/sales/:id", controllers.UpdateSale)
        kasirRoutes.DELETE("/sales/:id", controllers.DeleteSale)
    }

    // Pelanggan routes
    pelangganRoutes := router.Group("/pelanggan")
    pelangganRoutes.Use(middleware.AuthMiddleware(), middleware.CheckRoleMiddleware("pelanggan"))
    {
        pelangganRoutes.GET("/transactions", controllers.GetTransactions)
    }
}
