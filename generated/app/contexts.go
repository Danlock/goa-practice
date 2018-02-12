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
	uuid "github.com/satori/go.uuid"
	"net/http"
	"unicode/utf8"
)

// CreateAssetContext provides the asset create action context.
type CreateAssetContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreateAssetPayload
}

// NewCreateAssetContext parses the incoming request URL and body, performs validations and creates the
// context used by the asset controller create action.
func NewCreateAssetContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreateAssetContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreateAssetContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// createAssetPayload is the asset create action payload.
type createAssetPayload struct {
	// Type specific data
	Data *interface{} `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	// Name of asset
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Type of asset
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *createAssetPayload) Validate() (err error) {
	if payload.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	if payload.Type == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "type"))
	}
	if payload.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "data"))
	}
	if payload.Name != nil {
		if utf8.RuneCountInString(*payload.Name) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.name`, *payload.Name, utf8.RuneCountInString(*payload.Name), 3, true))
		}
	}
	if payload.Type != nil {
		if !(*payload.Type == "car") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`raw.type`, *payload.Type, []interface{}{"car"}))
		}
	}
	return
}

// Publicize creates CreateAssetPayload from createAssetPayload
func (payload *createAssetPayload) Publicize() *CreateAssetPayload {
	var pub CreateAssetPayload
	if payload.Data != nil {
		pub.Data = *payload.Data
	}
	if payload.Name != nil {
		pub.Name = *payload.Name
	}
	if payload.Type != nil {
		pub.Type = *payload.Type
	}
	return &pub
}

// CreateAssetPayload is the asset create action payload.
type CreateAssetPayload struct {
	// Type specific data
	Data interface{} `form:"data" json:"data" xml:"data"`
	// Name of asset
	Name string `form:"name" json:"name" xml:"name"`
	// Type of asset
	Type string `form:"type" json:"type" xml:"type"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateAssetPayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	if payload.Type == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "type"))
	}

	if utf8.RuneCountInString(payload.Name) < 3 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.name`, payload.Name, utf8.RuneCountInString(payload.Name), 3, true))
	}
	if !(payload.Type == "car") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`raw.type`, payload.Type, []interface{}{"car"}))
	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *CreateAssetContext) OK(r *Asset) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.asset+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// DeleteAssetContext provides the asset delete action context.
type DeleteAssetContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	AssetID string
}

// NewDeleteAssetContext parses the incoming request URL and body, performs validations and creates the
// context used by the asset controller delete action.
func NewDeleteAssetContext(ctx context.Context, r *http.Request, service *goa.Service) (*DeleteAssetContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := DeleteAssetContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramAssetID := req.Params["assetID"]
	if len(paramAssetID) > 0 {
		rawAssetID := paramAssetID[0]
		rctx.AssetID = rawAssetID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *DeleteAssetContext) OK(r *Asset) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.asset+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// ShowAssetContext provides the asset show action context.
type ShowAssetContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	AssetID uuid.UUID
}

// NewShowAssetContext parses the incoming request URL and body, performs validations and creates the
// context used by the asset controller show action.
func NewShowAssetContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowAssetContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowAssetContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramAssetID := req.Params["assetID"]
	if len(paramAssetID) > 0 {
		rawAssetID := paramAssetID[0]
		if assetID, err2 := uuid.FromString(rawAssetID); err2 == nil {
			rctx.AssetID = assetID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("assetID", rawAssetID, "uuid"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowAssetContext) OK(r *Asset) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.asset+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// ShowAllAssetContext provides the asset showAll action context.
type ShowAllAssetContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewShowAllAssetContext parses the incoming request URL and body, performs validations and creates the
// context used by the asset controller showAll action.
func NewShowAllAssetContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowAllAssetContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowAllAssetContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowAllAssetContext) OK(r AssetCollection) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.asset+json; type=collection")
	}
	if r == nil {
		r = AssetCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// UpdateAssetContext provides the asset update action context.
type UpdateAssetContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	AssetID uuid.UUID
	Payload *UpdateAssetPayload
}

// NewUpdateAssetContext parses the incoming request URL and body, performs validations and creates the
// context used by the asset controller update action.
func NewUpdateAssetContext(ctx context.Context, r *http.Request, service *goa.Service) (*UpdateAssetContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := UpdateAssetContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramAssetID := req.Params["assetID"]
	if len(paramAssetID) > 0 {
		rawAssetID := paramAssetID[0]
		if assetID, err2 := uuid.FromString(rawAssetID); err2 == nil {
			rctx.AssetID = assetID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("assetID", rawAssetID, "uuid"))
		}
	}
	return &rctx, err
}

// updateAssetPayload is the asset update action payload.
type updateAssetPayload struct {
	// Type specific data
	Data *interface{} `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	// Name of asset
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Type of asset
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *updateAssetPayload) Validate() (err error) {
	if payload.Name != nil {
		if utf8.RuneCountInString(*payload.Name) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.name`, *payload.Name, utf8.RuneCountInString(*payload.Name), 3, true))
		}
	}
	if payload.Type != nil {
		if !(*payload.Type == "car") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`raw.type`, *payload.Type, []interface{}{"car"}))
		}
	}
	return
}

// Publicize creates UpdateAssetPayload from updateAssetPayload
func (payload *updateAssetPayload) Publicize() *UpdateAssetPayload {
	var pub UpdateAssetPayload
	if payload.Data != nil {
		pub.Data = payload.Data
	}
	if payload.Name != nil {
		pub.Name = payload.Name
	}
	if payload.Type != nil {
		pub.Type = payload.Type
	}
	return &pub
}

// UpdateAssetPayload is the asset update action payload.
type UpdateAssetPayload struct {
	// Type specific data
	Data *interface{} `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	// Name of asset
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Type of asset
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *UpdateAssetPayload) Validate() (err error) {
	if payload.Name != nil {
		if utf8.RuneCountInString(*payload.Name) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.name`, *payload.Name, utf8.RuneCountInString(*payload.Name), 3, true))
		}
	}
	if payload.Type != nil {
		if !(*payload.Type == "car") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`raw.type`, *payload.Type, []interface{}{"car"}))
		}
	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *UpdateAssetContext) OK(r *Asset) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.asset+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// ShowStatusContext provides the status show action context.
type ShowStatusContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewShowStatusContext parses the incoming request URL and body, performs validations and creates the
// context used by the status controller show action.
func NewShowStatusContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowStatusContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowStatusContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowStatusContext) OK(r *Status) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.status+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}
