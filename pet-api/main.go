package main

import (
	"context"
	"pet-api/database"
	"pet-api/handlers"
	"pet-api/internal/telemetry"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func main() {
	shutdown := telemetry.InitTracer()
	defer shutdown(context.Background())

	database.InitMySQL()
	r := gin.Default()
	r.Use(otelgin.Middleware("pet-api"))

	r.POST("/pets", handlers.CreatePet)
	r.GET("/pets", handlers.ListPets)
	r.PUT("/pets/:id", handlers.UpdatePet)

	r.Run(":8080")
}
