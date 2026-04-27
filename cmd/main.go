package main

import (
	"context"
	"fmt"

	"github.com/Talan-Application/talan-back/internal/infrastructure/postgres"
)

func main() {
	ctx := context.Background()

	db, err := infrastructure_postgres.NewConnectionPool(ctx, infrastructure_postgres.NewConfigMust())
	if err != nil {
		fmt.Println("failed to connect to database: ", err)
	}
	defer db.Close()

}
