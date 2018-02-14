package design

import (
	"fmt"

	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/design/apidsl"
)

// addMetadata is needed to add bson struct tags for mgo
func addMetadata(name string) {
	apidsl.Metadata("struct:tag:bson", fmt.Sprintf("%s,omitempty", name))
	apidsl.Metadata("struct:tag:json", fmt.Sprintf("%s,omitempty", name))
	apidsl.Metadata("struct:tag:form", name)
}

var AssetType = apidsl.Type("AssetType", func() {
	apidsl.Attribute("_id", design.String, "Object ID attribute", func() {
		apidsl.Metadata("struct:field:type", "bson.ObjectId", "github.com/globalsign/mgo/bson")
		addMetadata("_id")
	})

	apidsl.Attribute("createdAt", design.DateTime, "timestamp of when the asset was created", func() {
		addMetadata("createdAt")
	})

	apidsl.Attribute("updatedAt", design.DateTime, "timestamp of when the asset was updated", func() {
		addMetadata("updatedAt")
	})

	apidsl.Attribute("name", design.String, "Name of asset", func() {
		addMetadata("name")
		apidsl.MinLength(3)
	})

	apidsl.Attribute("type", design.String, "Type of asset", func() {
		addMetadata("type")
		apidsl.Enum("car")
	})
	apidsl.Attribute("data", design.Any, "Type specific data", func() {
		addMetadata("data")
	})

})

var assetMedia = apidsl.MediaType("application/vnd.asset+json", func() {
	apidsl.Description("An asset can be pretty much anything")
	apidsl.Reference(AssetType)

	getAllAttributes := func() {
		apidsl.Attribute("_id")
		apidsl.Attribute("type")
		apidsl.Attribute("name")
		apidsl.Attribute("data")
		apidsl.Attribute("createdAt")
		apidsl.Attribute("updatedAt")
	}

	apidsl.Attributes(getAllAttributes)
	apidsl.View("default", getAllAttributes)
})

var _ = apidsl.Resource("asset", func() { // Resources group related API endpoints
	apidsl.BasePath("/asset") // together. They map to REST resources for REST

	apidsl.Action("showAll", func() {
		apidsl.Description("Get all assets")
		apidsl.Routing(apidsl.GET(""))
		apidsl.Response(design.OK, apidsl.CollectionOf(assetMedia))
	})

	apidsl.Action("create", func() {
		apidsl.Description("Create new assets")
		apidsl.Routing(apidsl.POST(""))
		apidsl.Payload(func() {
			apidsl.Reference(AssetType)
			apidsl.Attribute("type")
			apidsl.Attribute("name")
			apidsl.Attribute("data")
			apidsl.Required("name", "type", "data")
		})
		apidsl.Response(design.OK, assetMedia)
	})

	apidsl.Action("show", func() {
		apidsl.Description("Get specific asset")
		apidsl.Routing(apidsl.GET("/:assetID"))
		apidsl.Params(func() {
			apidsl.Param("assetID", design.String, "_id of an asset")
			apidsl.Required("assetID")
		})
		apidsl.Response(design.OK, assetMedia)
	})

	apidsl.Action("update", func() {
		apidsl.Description("Update specific asset")
		apidsl.Routing(apidsl.PUT("/:assetID"))
		apidsl.Params(func() {
			apidsl.Param("assetID", design.String, "_id of an asset")
			apidsl.Required("assetID")
		})

		apidsl.Payload(func() {
			apidsl.Reference(AssetType)
			apidsl.Attribute("type")
			apidsl.Attribute("name")
			apidsl.Attribute("data")
		})

		apidsl.Response(design.OK)
		apidsl.Response(design.NotFound)
	})

	apidsl.Action("delete", func() {
		apidsl.Description("Delete specific asset")
		apidsl.Routing(apidsl.DELETE("/:assetID"))
		apidsl.Params(func() {
			apidsl.Param("assetID", design.String, "_id of an asset")
			apidsl.Required("assetID")
		})
		apidsl.Response(design.OK)
	})
})
