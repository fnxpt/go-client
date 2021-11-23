// Code generated by go-swagger; DO NOT EDIT.

package replication

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
	"github.com/go-openapi/swag"
)

// NewStopReplicationParams creates a new StopReplicationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStopReplicationParams() *StopReplicationParams {
	return &StopReplicationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStopReplicationParamsWithTimeout creates a new StopReplicationParams object
// with the ability to set a timeout on a request.
func NewStopReplicationParamsWithTimeout(timeout time.Duration) *StopReplicationParams {
	return &StopReplicationParams{
		timeout: timeout,
	}
}

// NewStopReplicationParamsWithContext creates a new StopReplicationParams object
// with the ability to set a context for a request.
func NewStopReplicationParamsWithContext(ctx context.Context) *StopReplicationParams {
	return &StopReplicationParams{
		Context: ctx,
	}
}

// NewStopReplicationParamsWithHTTPClient creates a new StopReplicationParams object
// with the ability to set a custom HTTPClient for a request.
func NewStopReplicationParamsWithHTTPClient(client *http.Client) *StopReplicationParams {
	return &StopReplicationParams{
		HTTPClient: client,
	}
}

/* StopReplicationParams contains all the parameters to send to the API endpoint
   for the stop replication operation.

   Typically these are written to a http.Request.
*/
type StopReplicationParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* ID.

	   The ID of the execution.

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stop replication params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopReplicationParams) WithDefaults() *StopReplicationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stop replication params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopReplicationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stop replication params
func (o *StopReplicationParams) WithTimeout(timeout time.Duration) *StopReplicationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop replication params
func (o *StopReplicationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop replication params
func (o *StopReplicationParams) WithContext(ctx context.Context) *StopReplicationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop replication params
func (o *StopReplicationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop replication params
func (o *StopReplicationParams) WithHTTPClient(client *http.Client) *StopReplicationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop replication params
func (o *StopReplicationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the stop replication params
func (o *StopReplicationParams) WithXRequestID(xRequestID *string) *StopReplicationParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the stop replication params
func (o *StopReplicationParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the stop replication params
func (o *StopReplicationParams) WithID(id int64) *StopReplicationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the stop replication params
func (o *StopReplicationParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *StopReplicationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
