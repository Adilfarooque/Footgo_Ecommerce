package main

import (
	"fmt"
	"log"

	"github.com/Adilfarooque/Footgo_Ecommerce/config"
	"github.com/Adilfarooque/Footgo_Ecommerce/db"
	"github.com/Adilfarooque/Footgo_Ecommerce/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	cfig, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading the config file")
	}
	fmt.Println(cfig)
	db, err := db.ConnectDatabase(cfig)
	if err != nil {
		log.Fatalf("Error connecting to the database:%v", err)
	}
	r := gin.Default()
	//r.LoadHTMLGlob("template/*")
	userGroup := r.Group("/user")
	adminGroup := r.Group("/admin")
	routes.AdminRoutes(adminGroup, db)
	routes.UserRoutes(userGroup, db)

	listenAdder := fmt.Sprintf("%s:%s", cfig.DBPort, cfig.DBHost)
	fmt.Printf("Starting server on %s..\n", cfig.BASE_URL)
	if err := r.Run(cfig.BASE_URL); err != nil {
		log.Fatalf("Error starting server on %s:%s", listenAdder, err)
	}
}
