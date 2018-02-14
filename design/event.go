package design

import (
	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/design/apidsl"
)

var EventType = apidsl.Type("EvenType", func() {
	apidsl.Attribute("_id", design.String, "Object ID attribute", func() {
		apidsl.Metadata("struct:field:type", "bson.ObjectId", "github.com/globalsign/mgo/bson")
		AddMetadata("_id")
	})

	apidsl.Attribute("createdAt", design.DateTime, "timestamp of when the event was created", func() {
		AddMetadata("createdAt")
	})

	apidsl.Attribute("assetID", design.String, "Associated asset id", func() {
		apidsl.Metadata("struct:field:type", "bson.ObjectId", "github.com/globalsign/mgo/bson")
		AddMetadata("assetID")
	})

	apidsl.Attribute("type", design.String, "Type of event", func() {
		AddMetadata("type")
	})

	apidsl.Attribute("data", design.Any, "Type specific data", func() {
		AddMetadata("data")
	})

})

var EventMedia = apidsl.MediaType("application/vnd.event+json", func() {
	apidsl.Description("An event for an asset")
	apidsl.Reference(EventType)
	getAllAttributes := func() {
		apidsl.Attribute("_id")
		apidsl.Attribute("createdAt")
		apidsl.Attribute("assetID")
		apidsl.Attribute("type")
		apidsl.Attribute("data")
	}
	apidsl.Attributes(getAllAttributes)

	apidsl.View("default", getAllAttributes)

})

var _ = apidsl.Resource("event", func() {
	apidsl.Parent("asset")
	apidsl.BasePath("/event")
	apidsl.Action("showAll", func() {
		apidsl.Description("Get all events associated with an asset")
		apidsl.Routing(apidsl.GET("/all"))
		apidsl.Response(design.OK, apidsl.CollectionOf(EventMedia))
		apidsl.Response(design.NotFound)
	})

	apidsl.Action("publish", func() {
		apidsl.Description("Create new event and publish on the queue")
		apidsl.Routing(apidsl.POST("/publish"))
		apidsl.Payload(func() {
			apidsl.Reference(EventType)
			apidsl.Attribute("type")
			apidsl.Attribute("data")
			apidsl.Required("type", "data")
		})
		apidsl.Response(design.OK, EventMedia)
		apidsl.Response(design.NotFound)
	})

	apidsl.Action("subscribe", func() {
		apidsl.Description("Listen to all events for an asset")
		apidsl.Routing(apidsl.GET("/subscribe"))
		apidsl.Scheme("ws")
		apidsl.Response(design.SwitchingProtocols)
		apidsl.Response(design.NotFound)
	})
})
