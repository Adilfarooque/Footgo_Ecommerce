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
		products.POST("/filter", handlers.FilterCategory)
		products.GET("/image", handlers.ShowImages)
	}

	r.Use(middleware.UserAuthMiddleware())
	{

		address := r.Group("/address")
		{
			address.GET("", handlers.GetAllAddress)
			address.POST("", handlers.AddAddress)
			address.PUT("", handlers.UpdateAddress)
			address.DELETE("", handlers.DeleteAddressByID)
		}

		users := r.Group("/users")
		{
			users.GET("", handlers.UserDetails)
			users.PUT("", handlers.UpdateUserDetails)
			users.PUT("changepassword", handlers.ChangePassword)
		}

		wishlist := r.Group("/wishlist")
		{
			wishlist.POST("", handlers.AddToWishlist)
		}

		cart := r.Group("/cart")
		{
			cart.POST("", handlers.AddToCart)
			cart.DELETE("", handlers.RemoveFromCart)
			cart.GET("", handlers.DisplayCart)
			cart.DELETE("/empty", handlers.EmptyCart)
			cart.PUT("/updatequantityadd", handlers.UpdateQuantityAdd)
			cart.PUT("/updatequantityless", handlers.UpdateQuantityLess)
		}

		order := r.Group("/order")
		{
			order.GET("", handlers.GetOrderDetails)
			order.POST("", handlers.CancelOrder)
			order.GET("/place-order", handlers.PlaceOrderCOD)
		}

		checkout := r.Group("/checkout")
		{
			checkout.GET("",handlers.CheckOut)
			checkout.POST("",handlers.OrderItemsFromCart)
		}		
	}
	return r
}
