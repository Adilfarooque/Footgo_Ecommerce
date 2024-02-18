package routes

import (
	"github.com/Adilfarooque/Footgo_Ecommerce/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AdminRoutes(r *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {
	r.POST("adminlogin", handlers.LoginHandler)
	return r
}
