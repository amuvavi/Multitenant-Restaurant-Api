package routes

import (
	"theorcshack/api/handlers"
	"theorcshack/api/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.Use(middlewares.RateLimiter())

	// Public routes
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)

	// Protected review and rating routes (ensure X-Tenant-ID header is included in requests)
	router.POST("/dishes/:id/review", middlewares.AuthRequired(), handlers.ReviewAndRateDish)
	router.GET("/reviews/:id", middlewares.AuthRequired(), handlers.GetReview)
	router.PUT("/reviews/:id", middlewares.AuthRequired(), handlers.UpdateReview)
	router.DELETE("/reviews/:id", middlewares.AuthRequired(), handlers.DeleteReview)

	// Protected routes with tenant middleware
	api := router.Group("/api")
	api.Use(middlewares.AuthRequired())
	api.Use(middlewares.TenantMiddleware())
	{
		api.POST("/dishes", handlers.CreateDish)
		api.GET("/dishes", handlers.ListDishes)
		api.GET("/dishes/:id", handlers.GetDish)
		api.PUT("/dishes/:id", handlers.UpdateDish)
		api.DELETE("/dishes/:id", handlers.DeleteDish)
		api.GET("/search", handlers.SearchDishes)
	}
}
