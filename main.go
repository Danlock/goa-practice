//go:generate make

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/danlock/goa-practice/controller"
	"github.com/danlock/goa-practice/generated/app"
	"github.com/globalsign/mgo"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

//TODO: create users specific to application and use that instead of admin/default
func setupSessions() (*mgo.Session, *amqp.Connection) {
	mgoSession, err := mgo.Dial(
		fmt.Sprintf("%s:%s@localhost:27017", os.Getenv("MONGO_INITDB_ROOT_USERNAME"), os.Getenv("MONGO_INITDB_ROOT_PASSWORD")),
	)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB!\n%v", err)
	}

	amqpSession, err := amqp.Dial(
		fmt.Sprintf("amqp://%s:%s@localhost:5672", os.Getenv("RABBITMQ_DEFAULT_USER"), os.Getenv("RABBITMQ_DEFAULT_PASS")),
	)
	if err != nil {
		log.Fatalf("Error connecting to RabbitMQ!\n%v", err)
	}

	return mgoSession, amqpSession
}

func main() {
	//Load environment variables from docker dotenv, which will only exist in development
	_ = godotenv.Load("docker/.env")

	// Create service
	service := goa.New("practice")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	//Fails on use of websocket, must make custom logging middleware
	// service.Use(middleware.LogResponse())
	service.Use(middleware.ErrorHandler(service, true))
	//TODO: Recover middleware sends panics to the user as response, only mount in development.
	service.Use(middleware.Recover())

	mgoSession, amqpConn := setupSessions()

	//TODO: take cli flags and mount specific controllers
	app.MountStatusController(service, controller.NewStatusController(service, mgoSession))
	app.MountAssetController(service, controller.NewAssetController(service, mgoSession, amqpConn))
	app.MountDocumentationController(service, controller.NewDocumentationController(service))
	app.MountEventController(service, controller.NewEventController(service, mgoSession, amqpConn))
	app.MountSwaggerController(service, struct{ *goa.Controller }{Controller: service.NewController("SwaggerController")})

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
