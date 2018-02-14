package controller

import (
	"io"
	"log"
	"time"

	"github.com/danlock/goa-practice/generated/app"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/goadesign/goa"
	"github.com/streadway/amqp"
	"golang.org/x/net/websocket"
)

// EventController implements the event resource.
type EventController struct {
	*goa.Controller
	mgoSesh  *mgo.Session
	amqpConn *amqp.Connection
}

// NewEventController creates a event controller.
func NewEventController(service *goa.Service, mgoSesh *mgo.Session, amqpConn *amqp.Connection) *EventController {
	return &EventController{Controller: service.NewController("EventController"), mgoSesh: mgoSesh, amqpConn: amqpConn}
}

// Publish runs the publish action.
func (c *EventController) Publish(ctx *app.PublishEventContext) error {
	// EventController_Publish: start_implement
	if !bson.IsObjectIdHex(ctx.AssetID) {
		log.Printf("AssetID %s is not an ObjectID!", ctx.AssetID)
		return ctx.NotFound()
	}
	assetID := bson.ObjectIdHex(ctx.AssetID)

	now := time.Now()
	eventID := bson.NewObjectId()
	eventType := ctx.Payload.Type
	eventData := ctx.Payload.Data
	event := app.Event{
		ID:        &eventID,
		AssetID:   &assetID,
		Type:      &eventType,
		Data:      &eventData,
		CreatedAt: &now,
	}

	err := c.mgoSesh.DB("test").C("event").Insert(&event)
	if err != nil {
		log.Printf("Error inserting asset in mongo!\n%s", err)
		return err
	}

	return ctx.OK(&event)
	// EventController_Publish: end_implement
}

// ShowAll runs the showAll action.
func (c *EventController) ShowAll(ctx *app.ShowAllEventContext) error {
	// EventController_ShowAll: start_implement
	if !bson.IsObjectIdHex(ctx.AssetID) {
		log.Printf("AssetID %s is not an ObjectID!", ctx.AssetID)
		return ctx.NotFound()
	}
	assetID := bson.ObjectIdHex(ctx.AssetID)

	var events app.EventCollection
	err := c.mgoSesh.DB("test").C("event").Find(bson.M{"assetID": &assetID}).All(&events)
	if err != nil {
		log.Printf("Error finding all events from mongo!\n%s", err)
		return err
	}

	return ctx.OK(events)
	// EventController_ShowAll: end_implement
}

// Subscribe runs the subscribe action.
func (c *EventController) Subscribe(ctx *app.SubscribeEventContext) error {
	c.SubscribeWSHandler(ctx).ServeHTTP(ctx.ResponseWriter, ctx.Request)
	return nil
}

// SubscribeWSHandler establishes a websocket connection to run the subscribe action.
func (c *EventController) SubscribeWSHandler(ctx *app.SubscribeEventContext) websocket.Handler {
	return func(ws *websocket.Conn) {
		// EventController_Subscribe: start_implement

		// Put your logic here

		ws.Write([]byte("subscribe event"))
		// Dummy echo websocket server
		io.Copy(ws, ws)
		// EventController_Subscribe: end_implement
	}
}
