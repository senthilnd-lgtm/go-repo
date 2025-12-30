package main

import (
	"log"
	"movieapp/models"

	"movieapp/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.Use(gin.Logger())

	models.ConnectDatabase()

	movies := r.Group("/movies")
	{
		movies.POST("/", controllers.CreateMovie)
		movies.PUT("/:id", controllers.UpdateMovie)
		movies.DELETE("/:id", controllers.DeleteMovie)
		movies.DELETE("/", controllers.DeleteAllMovies)
		movies.GET("/", controllers.ListAllMovies)
		movies.GET("/one/:name", controllers.FindMovieByName)
		movies.GET("/all/:name", controllers.FindAllMoviesByName)
		movies.POST("/multiple", controllers.InsertMultipleMovies)
	}

	log.Println("Server Started")
	r.Run()

}
