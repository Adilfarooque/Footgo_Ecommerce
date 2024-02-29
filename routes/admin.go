package routes

import (
	"github.com/Adilfarooque/Footgo_Ecommerce/handlers"
	"github.com/Adilfarooque/Footgo_Ecommerce/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AdminRoutes(r *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {
	r.POST("adminlogin", handlers.LoginHandler)

	r.Use(middleware.AdminAuthMiddleware())
	{

		r.GET("/dashboard", handlers.DashBoard)
		r.GET("/sales-report",handlers.FilteredSalesReport)
		r.GET("/sales-report-date",handlers.SalesReportByDate)

		users := r.Group("/users")
		{
			users.GET("",handlers.GetUsers)
			users.PUT("/blocked",handlers.BlockedUser)
			users.PUT("/Unblock",handlers.UnBlockUser)
		}
	}
	return r
}

