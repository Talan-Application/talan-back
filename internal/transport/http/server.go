package http

import (
	"context"

	"github.com/Talan-Application/talan-back/internal/transport/http/handlers"
	"github.com/Talan-Application/talan-back/internal/transport/http/middlewares"
	"github.com/gin-gonic/gin"
)

type SimpleServer struct {
	srv         *gin.Engine
	userHandler *handlers.UserHandler
}

func NewHTTPServer(user UserService) *SimpleServer {
	userHandler := handlers.NewUserHandler(user)

	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())
	router.Use(middlewares.ErrorHandler())

	router.POST("/users", userHandler.CreateUser)
	router.GET("/users", userHandler.GetUsers)
	router.GET("/users/:id", userHandler.GetUser)
	router.PATCH("/users/:id", userHandler.UpdateUser)
	router.DELETE("/users/:id", userHandler.DeleteUser)

	protected := router.Group("/")
	protected.Use(middlewares.AuthMiddleware("jwt_secret"))
	{
		protected.GET("/ping", func(c *gin.Context) { c.JSON(200, gin.H{"message": "pong"}) })
	}
	return &SimpleServer{router, userHandler}
}

func (s *SimpleServer) Run(ctx context.Context) {
	err := s.srv.Run(":8080")
	if err != nil {
		panic(err)
	}
}
