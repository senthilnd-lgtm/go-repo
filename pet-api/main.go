package main

import (
	"context"
	"pet-api/database"
	"pet-api/handlers"
	"pet-api/internal/telemetry"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	ginprometheus "github.com/zsais/go-gin-prometheus"
)

func main() {
	shutdown := telemetry.InitTracer()
	defer shutdown(context.Background())

	database.InitMySQL()
	r := gin.Default()

	// Add OpenTelemetry middleware
	r.Use(otelgin.Middleware("pet-api"))

	// Add Prometheus middleware
	p := ginprometheus.NewPrometheus("pet_api")
	p.Use(r)

	r.POST("/pets", handlers.CreatePet)
	r.GET("/pets", handlers.ListPets)
	r.PUT("/pets/:id", handlers.UpdatePet)

	r.Run(":8080")
}
