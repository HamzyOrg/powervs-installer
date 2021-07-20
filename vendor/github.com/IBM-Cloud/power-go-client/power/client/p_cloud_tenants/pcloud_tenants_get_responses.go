// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudTenantsGetReader is a Reader for the PcloudTenantsGet structure.
type PcloudTenantsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudTenantsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPcloudTenantsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPcloudTenantsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPcloudTenantsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPcloudTenantsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudTenantsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudTenantsGetOK creates a PcloudTenantsGetOK with default headers values
func NewPcloudTenantsGetOK() *PcloudTenantsGetOK {
	return &PcloudTenantsGetOK{}
}

/*PcloudTenantsGetOK handles this case with default header values.

OK
*/
type PcloudTenantsGetOK struct {
	Payload *models.Tenant
}

func (o *PcloudTenantsGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/tenants/{tenant_id}][%d] pcloudTenantsGetOK  %+v", 200, o.Payload)
}

func (o *PcloudTenantsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tenant)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsGetBadRequest creates a PcloudTenantsGetBadRequest with default headers values
func NewPcloudTenantsGetBadRequest() *PcloudTenantsGetBadRequest {
	return &PcloudTenantsGetBadRequest{}
}

/*PcloudTenantsGetBadRequest handles this case with default header values.

Bad Request
*/
type PcloudTenantsGetBadRequest struct {
	Payload *models.Error
}

func (o *PcloudTenantsGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/tenants/{tenant_id}][%d] pcloudTenantsGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudTenantsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsGetForbidden creates a PcloudTenantsGetForbidden with default headers values
func NewPcloudTenantsGetForbidden() *PcloudTenantsGetForbidden {
	return &PcloudTenantsGetForbidden{}
}

/*PcloudTenantsGetForbidden handles this case with default header values.

Forbidden
*/
type PcloudTenantsGetForbidden struct {
	Payload *models.Error
}

func (o *PcloudTenantsGetForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/tenants/{tenant_id}][%d] pcloudTenantsGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudTenantsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsGetNotFound creates a PcloudTenantsGetNotFound with default headers values
func NewPcloudTenantsGetNotFound() *PcloudTenantsGetNotFound {
	return &PcloudTenantsGetNotFound{}
}

/*PcloudTenantsGetNotFound handles this case with default header values.

Not Found
*/
type PcloudTenantsGetNotFound struct {
	Payload *models.Error
}

func (o *PcloudTenantsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/tenants/{tenant_id}][%d] pcloudTenantsGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudTenantsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsGetInternalServerError creates a PcloudTenantsGetInternalServerError with default headers values
func NewPcloudTenantsGetInternalServerError() *PcloudTenantsGetInternalServerError {
	return &PcloudTenantsGetInternalServerError{}
}

/*PcloudTenantsGetInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudTenantsGetInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudTenantsGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/tenants/{tenant_id}][%d] pcloudTenantsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudTenantsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
