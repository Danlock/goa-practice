package design

import (
	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/design/apidsl"
)

var _ = apidsl.API("practice", func() {
	apidsl.Title("Practice for goa")
	apidsl.Description("Testing code and doc generation with goa")
	apidsl.Scheme("http")
	apidsl.Host("localhost:8080")
})

// CarMedia json blob
var CarMedia = apidsl.MediaType("application/practice.car+json", func() {
	apidsl.Description("Car with sensors")
	apidsl.Attributes(func() { // Attributes define the media type shape.
		apidsl.Attribute("id", design.Integer, "Unique car ID")
		apidsl.Attribute("name", design.String, "Name of car")
		apidsl.Attribute("plate", design.String, "License plate of car")
		apidsl.Required("id", "name")
	})
	apidsl.View("default", func() { // View defines a rendering of the media type.
		apidsl.Attribute("id") // Media types may have multiple views and must
		apidsl.Attribute("name")
	})
})

var _ = apidsl.Resource("car", func() { // Resources group related API endpoints
	apidsl.BasePath("/car")       // together. They map to REST resources for REST
	apidsl.DefaultMedia(CarMedia) // services.

	apidsl.Action("show", func() { // Actions define a single API endpoint together
		apidsl.Description("Get car by id")   // with its path, parameters (both path
		apidsl.Routing(apidsl.GET("/:carID")) // parameters and querystring values) and payload
		apidsl.Params(func() {                // (shape of the request body).
			apidsl.Param("carID", design.Integer, "Car ID")
		})
		apidsl.Response(design.OK)       // Responses define the shape and status code
		apidsl.Response(design.NotFound) // of HTTP responses.
	})

	apidsl.Action("create", func() { // Actions define a single API endpoint together
		apidsl.Description("Create car by id") // with its path, parameters (both path
		apidsl.Routing(apidsl.POST("/"))       // parameters and querystring values) and payload
		apidsl.Response(design.OK)             // Responses define the shape and status code
		apidsl.Response(design.NotFound)       // of HTTP responses.
	})
})

var _ = apidsl.Resource("officer", func() { // Resources group related API endpoints
	apidsl.BasePath("/officer")   // together. They map to REST resources for REST
	apidsl.DefaultMedia(CarMedia) // services.

	apidsl.Action("create", func() { // Actions define a single API endpoint together
		apidsl.Description("Create officer by id") // with its path, parameters (both path
		apidsl.Routing(apidsl.POST("/"))           // parameters and querystring values) and payload
		apidsl.Response(design.OK)                 // Responses define the shape and status code
		apidsl.Response(design.NotFound)           // of HTTP responses.
	})

	apidsl.Action("listen", func() { // Actions define a single API endpoint together
		apidsl.Description("Create officer by id") // with its path, parameters (both path
		apidsl.Routing(apidsl.GET("echo"))         // parameters and querystring values) and payload
		apidsl.Scheme("ws")
		apidsl.Response(design.OK)       // Responses define the shape and status code
		apidsl.Response(design.NotFound) // of HTTP responses.
	})
})
