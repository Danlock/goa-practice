package design

import (
	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/design/apidsl"
)

// StatusMedia json blob
var StatusMedia = apidsl.MediaType("application/vnd.status+json", func() {
	apidsl.Description("Status  of API and connections to remote services")
	apidsl.Attributes(func() { // Attributes define the media type shape.
		apidsl.Attribute("database", design.Boolean, "Is database connected and working?")
		apidsl.Required("database")
	})
	apidsl.View("default", func() { // View defines a rendering of the media type.
		apidsl.Attribute("database")
	})
})

var _ = apidsl.Resource("status", func() { // Resources group related API endpoints
	apidsl.BasePath("/status")       // together. They map to REST resources for REST
	apidsl.DefaultMedia(StatusMedia) // services.

	apidsl.Action("show", func() { // Actions define a single API endpoint together
		apidsl.Description("Get status of server and connections") // with its path, parameters (both path
		apidsl.Routing(apidsl.GET(""))                             // parameters and querystring values) and payload

		apidsl.Response(design.OK) // Responses define the shape and status code
	})
})
