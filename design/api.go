package design

import (
	"github.com/goadesign/goa/design/apidsl"
)

var _ = apidsl.API("practice", func() {
	apidsl.Title("Practice for goa")
	apidsl.Description("Testing code and doc generation with goa")
	apidsl.Scheme("http")
	apidsl.Host("localhost:8080")
})
