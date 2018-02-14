package design

import (
	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/design/apidsl"
)

var _ = apidsl.Resource("swagger", func() {
	apidsl.Origin("*", func() {
		apidsl.Methods("GET") // Allow all origins to retrieve the Swagger JSON (CORS)
	})
	apidsl.Files("/swagger.json", "generated/public/swagger/swagger.json")
})

var _ = apidsl.Resource("documentation", func() {
	apidsl.Origin("*", func() {
		apidsl.Methods("GET") // Allow all origins to retrieve the ReDoc Documentation
	})

	apidsl.Action("show", func() {
		apidsl.Description("ReDoc documentation generated from the swagger.json file")
		apidsl.Routing(apidsl.GET("/docs"))
		apidsl.Response(design.OK, "text/html; charset=utf-8")
	})
})
