package main

import (
	"fmt"
	"log"

	"github.com/Adilfarooque/Footgo_Ecommerce/config"
	"github.com/Adilfarooque/Footgo_Ecommerce/db"
	"github.com/Adilfarooque/Footgo_Ecommerce/routes"
	"github.com/gin-gonic/gin"
)

// @title Go + Gin Footgo E-Commerce API
// @version 1.0.0
// @description Footgo is an E-commerce platform to purchasing and selling shoes
// @contact.name API Support
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @host localhost:8000
// @BasePath /
// @query.collection.format multi

func main() {
	//Load configuration
	cfig, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading the config file:%v", err)
	}
	//Connect to the database
	db, err := db.ConnectDatabase(cfig)
	if err != nil {
		log.Fatal("Error connecting to the database:%v", err)
	}
	//Initialize Gin router
	r := gin.Default()
	//Load HTML templates
	r.LoadHTMLFiles("template/*")
	//Route Group for user and admin
	userGroup := r.Group("/user")
	//adminGroup := r.Group("/admin")
	routes.UserRoutes(userGroup, db)
	//routes.AdminRoutes(adminGroup,db)

	//Starting the server
	listenAdder := fmt.Sprintf("%s:%s", cfig.DBPort, cfig.DBHost)
	fmt.Printf("Starting server on %s..\n", cfig.BASE_URL)
	if err := r.Run(cfig.BASE_URL); err != nil {
		log.Fatalf("Error starting server on %s:%s", listenAdder, err)
	}
}
