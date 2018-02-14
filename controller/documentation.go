package controller

import (
	"github.com/danlock/goa-practice/generated/app"
	"github.com/goadesign/goa"
)

// DocumentationController implements the documentation resource.
type DocumentationController struct {
	*goa.Controller
}

// NewDocumentationController creates a documentation controller.
func NewDocumentationController(service *goa.Service) *DocumentationController {
	return &DocumentationController{Controller: service.NewController("DocumentationController")}
}

// Show runs the show action.
func (c *DocumentationController) Show(ctx *app.ShowDocumentationContext) error {
	// DocumentationController_Show: start_implement

	return ctx.OK([]byte(`
		<!DOCTYPE html>
		<html>
			<head>
				<title>Goa-Practice API Autogenerated Documentation</title>
				<!-- needed for adaptive design -->
				<meta charset="utf-8"/>
				<meta name="viewport" content="width=device-width, initial-scale=1">
		
				<!--
				ReDoc doesn't change outer page styles
				-->
				<style>
					body {
						margin: 0;
						padding: 0;
					}
				</style>
			</head>
			<body>
				<redoc spec-url='./swagger.json'></redoc>
				<script src="https://rebilly.github.io/ReDoc/releases/latest/redoc.min.js"> </script>
			</body>
		</html>
	`))
	// DocumentationController_Show: end_implement
}