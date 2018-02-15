package controller

import (
	"fmt"
	"log"
	"time"

	"encoding/json"

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
	amqpCh, err := c.amqpConn.Channel()
	if err != nil {
		log.Printf("Error opening channel to queue!\n%s", err)
		return err
	}
	defer amqpCh.Close()

	assetQ, err := amqpCh.QueueDeclare(fmt.Sprintf("asset:%s:event", ctx.AssetID), false, false, false, false, nil)
	if err != nil {
		log.Printf("Error declaring queue!\n%s", err)
		return err
	}

	eventJson, err := json.Marshal(&event)
	if err != nil {
		log.Printf("Error marshaling event JSON somehow!!\n%s", err)
		return err
	}

	err = amqpCh.Publish("", assetQ.Name, false, false, amqp.Publishing{
		ContentType: "apllication/json",
		Body:        eventJson,
	})
	if err != nil {
		log.Printf("Error publishing event!\n%s", err)
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

		amqpCh, err := c.amqpConn.Channel()
		if err != nil {
			log.Printf("Error opening channel to queue!\n%s", err)
			ws.WriteClose(1)
			return
		}
		defer amqpCh.Close()

		assetQ, err := amqpCh.QueueDeclare(fmt.Sprintf("asset:%s:event", ctx.AssetID), false, false, false, false, nil)
		if err != nil {
			log.Printf("Error declaring queue!\n%s", err)
			ws.WriteClose(2)
			return
		}
		ws.Write([]byte(fmt.Sprintf("subscribe event for asset %s", ctx.AssetID)))

		msgs, err := amqpCh.Consume(assetQ.Name, "", true, false, false, false, nil)
		if err != nil {
			log.Printf("Error consuming queue!\n%s", err)
			ws.WriteClose(3)
			return
		}

		for m := range msgs {
			// log.Printf("Received message!\n%s", string(m.Body))
			ws.Write(m.Body)
		}
		// EventController_Subscribe: end_implement
	}
}
