package handlers

import (
	"log"
	"net/http"
	"weight-tracker/pkg/api"

	"github.com/gin-gonic/gin"
)

type WeightHandler struct {
	WeightService api.WeightService
}

func NewWeightHandler(e *gin.Engine, weightService api.WeightService) {
	handler := &WeightHandler{
		WeightService: weightService,
	}

	e.POST("/api/v1/weight", handler.CreateWeightEntry)
}

func (handler *WeightHandler) CreateWeightEntry(c *gin.Context) {

	c.Header("Content-Type", "application/json")

	var newWeight api.NewWeightRequest
	err := c.ShouldBindJSON(&newWeight)

	if err != nil {
		log.Printf("handler error: %v", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	err = handler.WeightService.New(newWeight)

	if err != nil {
		log.Printf("service error: %v", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	response := map[string]string{
		"status": "success",
		"data":   "new create weight entry created",
	}

	c.JSON(http.StatusOK, response)
}
