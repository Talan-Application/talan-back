package main

import (
	"context"
	"fmt"

	"github.com/Talan-Application/talan-back/internal/infrastructure/postgres"
	"github.com/Talan-Application/talan-back/internal/repository"
	"github.com/Talan-Application/talan-back/internal/service"
	internalHttp "github.com/Talan-Application/talan-back/internal/transport/http"
)

func main() {
	ctx := context.Background()
	db, err := infrastructure_postgres.NewConnectionPool(ctx, infrastructure_postgres.NewConfigMust())
	if err != nil {
		fmt.Println("failed to connect to database: ", err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)

	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo)

	server := internalHttp.NewHTTPServer(userService, authService)
	server.Run(ctx)
}
