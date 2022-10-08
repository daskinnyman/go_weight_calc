package handlers

import (
	"log"
	"net/http"
	"weight-tracker/pkg/api"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService api.UserService
}

func NewUserHanlder(e *gin.Engine, userService api.UserService) {

	handler := &UserHandler{
		UserService: userService,
	}

	e.POST("/api/v1/user", handler.CreateUser)
}

func (handler *UserHandler) CreateUser(c *gin.Context) {

	c.Header("Content-Type", "application/json")

	var newUser api.NewUserRequest

	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		log.Printf("handler error: %v", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	err = handler.UserService.New(newUser)

	if err != nil {
		log.Printf("service error: %v", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	response := map[string]string{
		"status": "success",
		"data":   "new user created",
	}

	c.JSON(http.StatusOK, response)
}
