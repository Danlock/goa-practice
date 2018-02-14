package design

import (
	"fmt"

	"github.com/goadesign/goa/design/apidsl"
)

// addMetadata is needed to add bson struct tags for mgo
func AddMetadata(name string) {
	apidsl.Metadata("struct:tag:bson", fmt.Sprintf("%s,omitempty", name))
	apidsl.Metadata("struct:tag:json", fmt.Sprintf("%s,omitempty", name))
	apidsl.Metadata("struct:tag:form", name)
}

var _ = apidsl.API("practice", func() {
	apidsl.Title("Practice for goa")
	apidsl.Description("Testing code and doc generation with goa")
	apidsl.Scheme("http")
	apidsl.Host("localhost:8080")
})
