package http

import (
	"context"

	"github.com/Talan-Application/talan-back/internal/transport/http/handlers"
	"github.com/gin-gonic/gin"
)

type SimpleServer struct {
	srv         *gin.Engine
	userHandler *handlers.UserHandler
}

func NewHTTPServer(user UserService) *SimpleServer {
	userHandler := handlers.NewUserHandler(user)

	router := gin.Default()
	router.POST("/users", userHandler.CreateUser)
	router.GET("/users", userHandler.GetUsers)

	return &SimpleServer{router, userHandler}
}

func (s *SimpleServer) Run(ctx context.Context) {
	err := s.srv.Run(":8080")
	if err != nil {
		panic(err)
	}
}
