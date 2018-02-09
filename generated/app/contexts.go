// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "practice": Application Contexts
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
	"net/http"
	"strconv"
)

// CreateCarContext provides the car create action context.
type CreateCarContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewCreateCarContext parses the incoming request URL and body, performs validations and creates the
// context used by the car controller create action.
func NewCreateCarContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreateCarContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreateCarContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CreateCarContext) OK(r *PracticeCar) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/practice.car+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *CreateCarContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ShowCarContext provides the car show action context.
type ShowCarContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	CarID int
}

// NewShowCarContext parses the incoming request URL and body, performs validations and creates the
// context used by the car controller show action.
func NewShowCarContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowCarContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowCarContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramCarID := req.Params["carID"]
	if len(paramCarID) > 0 {
		rawCarID := paramCarID[0]
		if carID, err2 := strconv.Atoi(rawCarID); err2 == nil {
			rctx.CarID = carID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("carID", rawCarID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowCarContext) OK(r *PracticeCar) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/practice.car+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowCarContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// CreateOfficerContext provides the officer create action context.
type CreateOfficerContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewCreateOfficerContext parses the incoming request URL and body, performs validations and creates the
// context used by the officer controller create action.
func NewCreateOfficerContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreateOfficerContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreateOfficerContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CreateOfficerContext) OK(r *PracticeCar) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/practice.car+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *CreateOfficerContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListenOfficerContext provides the officer listen action context.
type ListenOfficerContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListenOfficerContext parses the incoming request URL and body, performs validations and creates the
// context used by the officer controller listen action.
func NewListenOfficerContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListenOfficerContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListenOfficerContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListenOfficerContext) OK(r *PracticeCar) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/practice.car+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListenOfficerContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}