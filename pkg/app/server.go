package app

import (
	"log"
	"weight-tracker/pkg/api"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router        *gin.Engine
	userService   api.UserService
	weightService api.WeightService
}

func NewServer(router *gin.Engine, userService api.UserService, weightService api.WeightService) *Server {
	return &Server{
		router:        router,
		userService:   userService,
		weightService: weightService,
	}
}

func (s *Server) Run() error {
	r := s.Routes()
	err := r.Run(":5001")

	if err != nil {
		log.Printf("Server - there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}
