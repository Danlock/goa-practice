package controller

import (
	"github.com/danlock/goa-practice/generated/app"
	"github.com/goadesign/goa"
	"golang.org/x/net/websocket"
	"io"
)

// OfficerController implements the officer resource.
type OfficerController struct {
	*goa.Controller
}

// NewOfficerController creates a officer controller.
func NewOfficerController(service *goa.Service) *OfficerController {
	return &OfficerController{Controller: service.NewController("OfficerController")}
}

// Create runs the create action.
func (c *OfficerController) Create(ctx *app.CreateOfficerContext) error {
	// OfficerController_Create: start_implement

	// Put your logic here

	res := &app.PracticeCar{}
	return ctx.OK(res)

	// OfficerController_Create: end_implement
}

// Listen runs the listen action.
func (c *OfficerController) Listen(ctx *app.ListenOfficerContext) error {
	c.ListenWSHandler(ctx).ServeHTTP(ctx.ResponseWriter, ctx.Request)
	return nil
}

// ListenWSHandler establishes a websocket connection to run the listen action.
func (c *OfficerController) ListenWSHandler(ctx *app.ListenOfficerContext) websocket.Handler {
	return func(ws *websocket.Conn) {
		// OfficerController_Listen: start_implement

		// Put your logic here

		ws.Write([]byte("listen officer"))
		// Dummy echo websocket server
		io.Copy(ws, ws)
		// OfficerController_Listen: end_implement
	}
}
