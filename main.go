package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go-fiber-rest/api/routes"
	"go-fiber-rest/pkg/product"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func main() {
	db, err := databaseConnection()
	l := log.New(os.Stdout, "[products-api] ", log.LstdFlags)
	if err != nil {
		l.Fatal("Database Connection Error $s", err)
	}

	l.Println("server: Connected to database!")

	prodCollection := db.Collection("products")
	prodRepo := product.NewRepo(prodCollection)
	prodService := product.NewService(prodRepo)

	app := fiber.New()
	api := app.Group("/api")
	routes.ProductRouter(api, prodService)

	err = app.Listen(":3000")
	if err != nil {
		println("Error trying to start the server")
		os.Exit(1)
	}
}

func databaseConnection() (*mongo.Database, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb://localhost:27017/fiber",
	).SetServerSelectionTimeout(5*time.Second))

	if err != nil {
		return nil, err
	}

	db := client.Database("products")
	return db, nil
}
