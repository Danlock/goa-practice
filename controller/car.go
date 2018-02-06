package controller

import (
	"github.com/danlock/goa-practice/generated/app"
	"github.com/goadesign/goa"
)

// CarController implements the car resource.
type CarController struct {
	*goa.Controller
}

// NewCarController creates a car controller.
func NewCarController(service *goa.Service) *CarController {
	return &CarController{Controller: service.NewController("CarController")}
}

// Create runs the create action.
func (c *CarController) Create(ctx *app.CreateCarContext) error {
	// CarController_Create: start_implement

	// Put your logic here

	res := &app.PracticeCar{}
	return ctx.OK(res)

	// CarController_Create: end_implement
}

// Show runs the show action.
func (c *CarController) Show(ctx *app.ShowCarContext) error {
	// CarController_Show: start_implement

	// Put your logic here

	res := &app.PracticeCar{}
	return ctx.OK(res)

	// CarController_Show: end_implement
}
