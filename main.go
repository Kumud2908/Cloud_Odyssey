package main

import (
	"golang-hospital-management/database"
	"golang-hospital-management/middleware"
	"golang-hospital-management/routes"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var AppointmentConnection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	
	routes.AppointmentRoutes(router)

	
	router.Run(":" + port)

}
