package main

import (
	"database/sql"
	"fmt"
	"os"
	"weight-tracker/pkg/api"
	"weight-tracker/pkg/handlers"
	"weight-tracker/pkg/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup error: %s\\n", err)
		os.Exit(1)
	}
}

func run() error {
	connectionString := "postgres://postgres:example@127.0.0.1:5432/postgres?sslmode=disable"
	db, err := setupDatabase(connectionString)

	if err != nil {
		return err
	}

	storage := repository.NewStorage(db)

	err = storage.RunMigrations("postgres://postgres:example@127.0.0.1:5432/postgres?sslmode=disable")

	if err != nil {
		return err
	}

	router := gin.Default()
	router.Use(cors.Default())

	// Inject repo into services
	userService := api.NewUserService(storage)
	weightService := api.NewWeightService(storage)

	// Inject services into routes
	handlers.NewUserHanlder(router, userService)
	handlers.NewWeightHandler(router, weightService)

	err = router.Run(":5001")
	if err != nil {
		fmt.Printf("Server - there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}

func setupDatabase(connString string) (*sql.DB, error) {
	// change "postgres" for whatever supported database you want to use
	db, err := sql.Open("postgres", connString)

	if err != nil {
		return nil, err
	}

	// ping the DB to ensure that it is connected
	err = db.Ping()

	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully created connection to database")

	return db, nil
}
