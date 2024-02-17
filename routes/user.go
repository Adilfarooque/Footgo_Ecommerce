package routes

import (
	"github.com/Adilfarooque/Footgo_Ecommerce/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {
	r.POST("/signup", handlers.UserSignup)
	r.POST("/login", handlers.Userlogin)
	return r
}
