//go:generate make

package main

import (
	"github.com/danlock/goa-practice/controller"
	"github.com/danlock/goa-practice/generated/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("practice")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "car" controller
	c := controller.NewCarController(service)
	app.MountCarController(service, c)
	// Mount "officer" controller
	c2 := controller.NewOfficerController(service)
	app.MountOfficerController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
