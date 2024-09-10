package main

import (
	"context"
	"flag"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/kaiquecaires/hotel_reservation/api"
	"github.com/kaiquecaires/hotel_reservation/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbUri = "mongodb://localhost:27017"
const dbName = "hotel-reservation"
const userCollection = "users"

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))

	if err != nil {
		log.Fatal(err)
	}

	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the API server")
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c fiber.Ctx, err error) error {
			return c.JSON(map[string]string{"error": err.Error()})
		},
	})

	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(map[string]string{"msg": "Server is on fire!"})
	})
	apiv1 := app.Group("/api/v1")
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))
	apiv1.Post("/user", userHandler.HandlePostUser)
	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)
	apiv1.Delete("/user/:id", userHandler.HandleDeleteUser)
	apiv1.Put("/user/:id", userHandler.HandlePutUser)
	app.Listen(*listenAddr)
}
