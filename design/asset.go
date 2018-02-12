package design

import (
	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/design/apidsl"
)

func immutableAssetAttributes(modifiers ...func()) {
	apidsl.Attribute("_id", design.String, "Identifier for asset", func() {
		//Add tags so that mgo treats this as the primary key field instead of adding a new one
		apidsl.Metadata("struct:field:type", "bson.ObjectId", "github.com/globalsign/mgo/bson")
		apidsl.Metadata("struct:tag:bson", "_id")
		apidsl.Metadata("struct:tag:json", "_id")
		apidsl.Metadata("struct:tag:form", "_id")
	})
	apidsl.Attribute("createdAt", design.DateTime, "timestamp of when the asset was created")
	apidsl.Attribute("updatedAt", design.DateTime, "timestamp of when the asset was updated")

	for _, f := range modifiers {
		f()
	}
}

func mutableAssetAttributes(modifiers ...func()) {
	apidsl.Attribute("name", design.String, "Name of asset", func() {
		apidsl.MinLength(3)
	})
	apidsl.Attribute("type", design.String, "Type of asset", func() {
		apidsl.Enum("car")
	})
	apidsl.Attribute("data", design.Any, "Type specific data")

	for _, f := range modifiers {
		f()
	}
}

var AssetMedia = apidsl.MediaType("application/vnd.asset+json", func() {
	apidsl.Description("An asset can be pretty much anything")
	apidsl.Attributes(func() {
		mutableAssetAttributes()
		immutableAssetAttributes()
		apidsl.Required("_id", "name", "type", "data")
	})

	apidsl.View("default", func() { // View defines a rendering of the media type.
		apidsl.Attribute("_id")
		apidsl.Attribute("type")
		apidsl.Attribute("name")
		apidsl.Attribute("data")
		apidsl.Attribute("createdAt")
		apidsl.Attribute("updatedAt")
	})

})

var _ = apidsl.Resource("asset", func() { // Resources group related API endpoints
	apidsl.BasePath("/asset")       // together. They map to REST resources for REST
	apidsl.DefaultMedia(AssetMedia) // services.

	apidsl.Action("showAll", func() {
		apidsl.Description("Get all assets")
		apidsl.Routing(apidsl.GET(""))
		apidsl.Response(design.OK, apidsl.CollectionOf(AssetMedia))
	})

	apidsl.Action("create", func() {
		apidsl.Description("Create new assets")
		apidsl.Routing(apidsl.POST(""))
		apidsl.Payload(func() {
			mutableAssetAttributes()
			apidsl.Required("name", "type", "data")
		})
		apidsl.Response(design.OK)
	})

	apidsl.Action("show", func() {
		apidsl.Description("Get specific asset")
		apidsl.Routing(apidsl.GET("/:assetID"))
		apidsl.Params(func() {
			apidsl.Param("assetID", design.UUID, "Asset ID")
			apidsl.Required("assetID")
		})
		apidsl.Response(design.OK)
	})

	apidsl.Action("update", func() {
		apidsl.Description("Update specific asset")
		apidsl.Routing(apidsl.PUT("/:assetID"))
		apidsl.Params(func() {
			apidsl.Param("assetID", design.UUID, "Asset ID")
			apidsl.Required("assetID")
		})

		apidsl.Payload(func() {
			mutableAssetAttributes()
		})

		apidsl.Response(design.OK)
	})

	apidsl.Action("delete", func() {
		apidsl.Description("Delete specific asset")
		apidsl.Routing(apidsl.DELETE("/:assetID"))

		apidsl.Response(design.OK)
	})
})
