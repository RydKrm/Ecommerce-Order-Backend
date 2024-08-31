package main

import (
	"delivery/database"
	"delivery/modules"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load env file
	err := godotenv.Load();
	if err!= nil{
		log.Fatalf("Error loading .env file");
	}
   // Create Database connection
	database.Connection()

	// Disable gin's debug mode
	if os.Getenv("GIN_MODE") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
    //  Initialize the router
	router := gin.Default();
	// Check for default route
	router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status":  true,
            "message": "Server working",
        })
    })
	// Initialize all routes
	modules.RouteList(router);

	// Set up the trusted proxy
	router.SetTrustedProxies([]string{"127.0.01"});

	// Find the port
	port  := os.Getenv("PORT");

	if port == ""{
		port = "5000"
	}

	// finally run the server
	if err := router.Run(":"+port); err!=nil{
		log.Fatalf("Field to run the server : %v \n", err)
	} else {
		fmt.Println("Server running on port http://localhost:"+port)
	}



}