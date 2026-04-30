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
	authHandler *handlers.AuthHandler
}

func NewHTTPServer(user UserService, auth AuthService) *SimpleServer {
	userHandler := handlers.NewUserHandler(user)
	authHandler := handlers.NewAuthHandler(auth)

	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())
	router.Use(middlewares.ErrorHandler())

	router.POST("/signup", authHandler.Registration)
	router.POST("/login")

	protected := router.Group("/")
	protected.Use(middlewares.AuthMiddleware("jwt_secret"))
	{
		protected.GET("/ping", func(c *gin.Context) { c.JSON(200, gin.H{"message": "pong"}) })
		protected.POST("/users", userHandler.CreateUser)
		protected.GET("/users", userHandler.GetUsers)
		protected.GET("/users/:id", userHandler.GetUser)
		protected.PATCH("/users/:id", userHandler.UpdateUser)
		protected.DELETE("/users/:id", userHandler.DeleteUser)
	}

	return &SimpleServer{
		router,
		userHandler,
		authHandler,
	}
}

func (s *SimpleServer) Run(ctx context.Context) {
	err := s.srv.Run(":8080")
	if err != nil {
		panic(err)
	}
}
