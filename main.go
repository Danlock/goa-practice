//go:generate make

package main

import (
	"log"

	"github.com/danlock/goa-practice/controller"
	"github.com/danlock/goa-practice/generated/app"
	"github.com/globalsign/mgo"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func setupConnectionPools() (*mgo.Session, *amqp.Connection) {
	mgoSession, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatalf("Error connecting to MongoDB!")
	}

	amqpSession, err := amqp.Dial("amqp://${RABBITMQ_DEFAULT_USER}:${RABBITMQ_DEFAULT_PASS}@localhost:5672/")
	if err != nil {
		log.Fatalf("Error connecting to RabbitMQ!")
	}

	return mgoSession, amqpSession
}

func main() {
	_ = godotenv.Load("docker/.env")

	// Create service
	service := goa.New("practice")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	app.MountStatusController(service, controller.NewStatusController(service))

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
