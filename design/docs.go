package design

import (
	"github.com/goadesign/goa/design/apidsl"
)

var _ = apidsl.Resource("swagger", func() {
	apidsl.Origin("*", func() {
		apidsl.Methods("GET") // Allow all origins to retrieve the Swagger JSON (CORS)
	})
	apidsl.Files("/swagger.json", "generated/public/swagger/swagger.json")
})
