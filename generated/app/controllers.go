// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "practice": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/danlock/goa-practice/design
// --out=$(GOPATH)/src/github.com/danlock/goa-practice/generated
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// AssetController is the controller interface for the Asset actions.
type AssetController interface {
	goa.Muxer
	Create(*CreateAssetContext) error
	Delete(*DeleteAssetContext) error
	Show(*ShowAssetContext) error
	ShowAll(*ShowAllAssetContext) error
	Update(*UpdateAssetContext) error
}

// MountAssetController "mounts" a Asset resource controller on the given service.
func MountAssetController(service *goa.Service, ctrl AssetController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateAssetContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateAssetPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/asset", ctrl.MuxHandler("create", h, unmarshalCreateAssetPayload))
	service.LogInfo("mount", "ctrl", "Asset", "action", "Create", "route", "POST /asset")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteAssetContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	service.Mux.Handle("DELETE", "/asset/:assetID", ctrl.MuxHandler("delete", h, nil))
	service.LogInfo("mount", "ctrl", "Asset", "action", "Delete", "route", "DELETE /asset/:assetID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowAssetContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/asset/:assetID", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Asset", "action", "Show", "route", "GET /asset/:assetID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowAllAssetContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.ShowAll(rctx)
	}
	service.Mux.Handle("GET", "/asset", ctrl.MuxHandler("showAll", h, nil))
	service.LogInfo("mount", "ctrl", "Asset", "action", "ShowAll", "route", "GET /asset")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateAssetContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateAssetPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	service.Mux.Handle("PUT", "/asset/:assetID", ctrl.MuxHandler("update", h, unmarshalUpdateAssetPayload))
	service.LogInfo("mount", "ctrl", "Asset", "action", "Update", "route", "PUT /asset/:assetID")
}

// unmarshalCreateAssetPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateAssetPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createAssetPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateAssetPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateAssetPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &updateAssetPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// StatusController is the controller interface for the Status actions.
type StatusController interface {
	goa.Muxer
	Show(*ShowStatusContext) error
}

// MountStatusController "mounts" a Status resource controller on the given service.
func MountStatusController(service *goa.Service, ctrl StatusController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowStatusContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/status", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Status", "action", "Show", "route", "GET /status")
}

// SwaggerController is the controller interface for the Swagger actions.
type SwaggerController interface {
	goa.Muxer
	goa.FileServer
}

// MountSwaggerController "mounts" a Swagger resource controller on the given service.
func MountSwaggerController(service *goa.Service, ctrl SwaggerController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/swagger.json", ctrl.MuxHandler("preflight", handleSwaggerOrigin(cors.HandlePreflight()), nil))

	h = ctrl.FileHandler("/swagger.json", "generated/public/swagger/swagger.json")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger.json", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "generated/public/swagger/swagger.json", "route", "GET /swagger.json")
}

// handleSwaggerOrigin applies the CORS response headers corresponding to the origin.
func handleSwaggerOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}
