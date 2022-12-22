// Code generated by go-swagger; DO NOT EDIT.

package metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/openziti-test-kitchen/zrok/rest_model_zrok"
)

// GetEnvironmentDetailHandlerFunc turns a function with the right signature into a get environment detail handler
type GetEnvironmentDetailHandlerFunc func(GetEnvironmentDetailParams, *rest_model_zrok.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetEnvironmentDetailHandlerFunc) Handle(params GetEnvironmentDetailParams, principal *rest_model_zrok.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetEnvironmentDetailHandler interface for that can handle valid get environment detail params
type GetEnvironmentDetailHandler interface {
	Handle(GetEnvironmentDetailParams, *rest_model_zrok.Principal) middleware.Responder
}

// NewGetEnvironmentDetail creates a new http.Handler for the get environment detail operation
func NewGetEnvironmentDetail(ctx *middleware.Context, handler GetEnvironmentDetailHandler) *GetEnvironmentDetail {
	return &GetEnvironmentDetail{Context: ctx, Handler: handler}
}

/*
	GetEnvironmentDetail swagger:route GET /detail/environment metadata getEnvironmentDetail

GetEnvironmentDetail get environment detail API
*/
type GetEnvironmentDetail struct {
	Context *middleware.Context
	Handler GetEnvironmentDetailHandler
}

func (o *GetEnvironmentDetail) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetEnvironmentDetailParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *rest_model_zrok.Principal
	if uprinc != nil {
		principal = uprinc.(*rest_model_zrok.Principal) // this is really a rest_model_zrok.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetEnvironmentDetailBody get environment detail body
//
// swagger:model GetEnvironmentDetailBody
type GetEnvironmentDetailBody struct {

	// env z Id
	EnvZID string `json:"envZId,omitempty"`
}

// Validate validates this get environment detail body
func (o *GetEnvironmentDetailBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get environment detail body based on context it is used
func (o *GetEnvironmentDetailBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetEnvironmentDetailBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetEnvironmentDetailBody) UnmarshalBinary(b []byte) error {
	var res GetEnvironmentDetailBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
