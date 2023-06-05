package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	controller "github.com/sanmathi-g/restapi/controllers"
	"github.com/sanmathi-g/restapi/database"
)

func main() {
	fmt.Println("Starting application ...")
	database.DatabaseConnection()

	router := gin.Default()
	router.GET("/movies/:id", controller.ReadMovie)
	router.GET("/movies", controller.Readmovies)
	router.POST("/movies", controller.CreateMovie)
	router.PUT("/moviess/:id", controller.UpdateMovie)
	router.DELETE("/moviess/:id", controller.DeleteMovie)
	router.Run(":5000")
}
