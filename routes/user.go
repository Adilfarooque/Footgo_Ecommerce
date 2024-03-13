package routes

import (
	"github.com/Adilfarooque/Footgo_Ecommerce/handlers"
	"github.com/Adilfarooque/Footgo_Ecommerce/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {
	r.POST("/signup", handlers.UserSignup)
	r.POST("/login", handlers.Userlogin)

	products := r.Group("/products")
	{
		products.GET("", handlers.ShowAllProducts)
		products.POST("/filter",handlers.FilterCategory)
		products.GET("/image",handlers.ShowImages)
	}

	r.Use(middleware.UserAuthMiddleware())
	{
		address := r.Group("/address")
		{
			address.GET("",handlers.GetAllAddress)
			address.POST("",handlers.AddAddress)
			
		}
	}
	return r
}
