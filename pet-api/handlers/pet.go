package handlers

import (
	"log"
	"net/http"
	"pet-api/database"
	"pet-api/models"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func CreatePet(c *gin.Context) {
	ctx := c.Request.Context()
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("handler", "CreatePet"))

	traceID := span.SpanContext().TraceID().String()
	log.Printf("CreatePet TraceID: %s", traceID)

	var pet models.Pet
	if err := c.ShouldBindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.WithContext(c).Create(&pet)
	c.JSON(http.StatusCreated, pet)
}

func ListPets(c *gin.Context) {

	ctx := c.Request.Context()
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("handler", "ListPet"))

	traceID := span.SpanContext().TraceID().String()
	log.Printf("ListPet TraceID: %s", traceID)

	var pets []models.Pet
	database.DB.WithContext(c).Find(&pets)
	c.JSON(http.StatusOK, pets)
}

func UpdatePet(c *gin.Context) {

	ctx := c.Request.Context()
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("handler", "Updatepet"))

	traceID := span.SpanContext().TraceID().String()
	log.Printf("Updatepet TraceID: %s", traceID)

	var pet models.Pet
	id := c.Param("id")

	if err := database.DB.WithContext(c).First(&pet, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pet not found"})
		return
	}

	if err := c.ShouldBindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.WithContext(c).Save(&pet)
	c.JSON(http.StatusOK, pet)
}
