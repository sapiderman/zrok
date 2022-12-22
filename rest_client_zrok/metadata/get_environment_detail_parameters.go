// Code generated by go-swagger; DO NOT EDIT.

package metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetEnvironmentDetailParams creates a new GetEnvironmentDetailParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEnvironmentDetailParams() *GetEnvironmentDetailParams {
	return &GetEnvironmentDetailParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEnvironmentDetailParamsWithTimeout creates a new GetEnvironmentDetailParams object
// with the ability to set a timeout on a request.
func NewGetEnvironmentDetailParamsWithTimeout(timeout time.Duration) *GetEnvironmentDetailParams {
	return &GetEnvironmentDetailParams{
		timeout: timeout,
	}
}

// NewGetEnvironmentDetailParamsWithContext creates a new GetEnvironmentDetailParams object
// with the ability to set a context for a request.
func NewGetEnvironmentDetailParamsWithContext(ctx context.Context) *GetEnvironmentDetailParams {
	return &GetEnvironmentDetailParams{
		Context: ctx,
	}
}

// NewGetEnvironmentDetailParamsWithHTTPClient creates a new GetEnvironmentDetailParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEnvironmentDetailParamsWithHTTPClient(client *http.Client) *GetEnvironmentDetailParams {
	return &GetEnvironmentDetailParams{
		HTTPClient: client,
	}
}

/*
GetEnvironmentDetailParams contains all the parameters to send to the API endpoint

	for the get environment detail operation.

	Typically these are written to a http.Request.
*/
type GetEnvironmentDetailParams struct {

	// Body.
	Body GetEnvironmentDetailBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get environment detail params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEnvironmentDetailParams) WithDefaults() *GetEnvironmentDetailParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get environment detail params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEnvironmentDetailParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get environment detail params
func (o *GetEnvironmentDetailParams) WithTimeout(timeout time.Duration) *GetEnvironmentDetailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get environment detail params
func (o *GetEnvironmentDetailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get environment detail params
func (o *GetEnvironmentDetailParams) WithContext(ctx context.Context) *GetEnvironmentDetailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get environment detail params
func (o *GetEnvironmentDetailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get environment detail params
func (o *GetEnvironmentDetailParams) WithHTTPClient(client *http.Client) *GetEnvironmentDetailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get environment detail params
func (o *GetEnvironmentDetailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get environment detail params
func (o *GetEnvironmentDetailParams) WithBody(body GetEnvironmentDetailBody) *GetEnvironmentDetailParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get environment detail params
func (o *GetEnvironmentDetailParams) SetBody(body GetEnvironmentDetailBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetEnvironmentDetailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
