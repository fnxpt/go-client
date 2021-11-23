// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

// GetProjectDeletableReader is a Reader for the GetProjectDeletable structure.
type GetProjectDeletableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectDeletableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectDeletableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetProjectDeletableUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProjectDeletableForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProjectDeletableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProjectDeletableInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProjectDeletableOK creates a GetProjectDeletableOK with default headers values
func NewGetProjectDeletableOK() *GetProjectDeletableOK {
	return &GetProjectDeletableOK{}
}

/* GetProjectDeletableOK describes a response with status code 200, with default header values.

Return deletable status of the project.
*/
type GetProjectDeletableOK struct {
	Payload *models.ProjectDeletable
}

func (o *GetProjectDeletableOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/_deletable][%d] getProjectDeletableOK  %+v", 200, o.Payload)
}
func (o *GetProjectDeletableOK) GetPayload() *models.ProjectDeletable {
	return o.Payload
}

func (o *GetProjectDeletableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectDeletable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectDeletableUnauthorized creates a GetProjectDeletableUnauthorized with default headers values
func NewGetProjectDeletableUnauthorized() *GetProjectDeletableUnauthorized {
	return &GetProjectDeletableUnauthorized{}
}

/* GetProjectDeletableUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetProjectDeletableUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetProjectDeletableUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/_deletable][%d] getProjectDeletableUnauthorized  %+v", 401, o.Payload)
}
func (o *GetProjectDeletableUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetProjectDeletableUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectDeletableForbidden creates a GetProjectDeletableForbidden with default headers values
func NewGetProjectDeletableForbidden() *GetProjectDeletableForbidden {
	return &GetProjectDeletableForbidden{}
}

/* GetProjectDeletableForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetProjectDeletableForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetProjectDeletableForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/_deletable][%d] getProjectDeletableForbidden  %+v", 403, o.Payload)
}
func (o *GetProjectDeletableForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetProjectDeletableForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectDeletableNotFound creates a GetProjectDeletableNotFound with default headers values
func NewGetProjectDeletableNotFound() *GetProjectDeletableNotFound {
	return &GetProjectDeletableNotFound{}
}

/* GetProjectDeletableNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetProjectDeletableNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetProjectDeletableNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/_deletable][%d] getProjectDeletableNotFound  %+v", 404, o.Payload)
}
func (o *GetProjectDeletableNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetProjectDeletableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectDeletableInternalServerError creates a GetProjectDeletableInternalServerError with default headers values
func NewGetProjectDeletableInternalServerError() *GetProjectDeletableInternalServerError {
	return &GetProjectDeletableInternalServerError{}
}

/* GetProjectDeletableInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetProjectDeletableInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *GetProjectDeletableInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/_deletable][%d] getProjectDeletableInternalServerError  %+v", 500, o.Payload)
}
func (o *GetProjectDeletableInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetProjectDeletableInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
