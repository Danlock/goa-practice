package controller

import (
	"log"

	"github.com/danlock/goa-practice/generated/app"
	"github.com/globalsign/mgo"
	"github.com/goadesign/goa"
)

// StatusController implements the status resource.
type StatusController struct {
	*goa.Controller
	mgoSesh *mgo.Session
}

// NewStatusController creates a status controller.
func NewStatusController(service *goa.Service, mgoSesh *mgo.Session) *StatusController {
	return &StatusController{Controller: service.NewController("StatusController"), mgoSesh: mgoSesh}
}

// Show runs the show action.
func (c *StatusController) Show(ctx *app.ShowStatusContext) error {
	// StatusController_Show: start_implement

	mgoStatus := false

	servers := c.mgoSesh.LiveServers()
	log.Printf("Currently connected to MongoDB servers : %s", servers)
	if len(servers) > 0 {
		mgoStatus = true
	}

	res := &app.Status{Database: mgoStatus}
	return ctx.OK(res)
	// StatusController_Show: end_implement
}
