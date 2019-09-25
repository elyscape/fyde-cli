// Code generated by go-swagger; DO NOT EDIT.

package access_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/oNaiPs/fyde-cli/models"
)

// DeletePolicyReader is a Reader for the DeletePolicy structure.
type DeletePolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeletePolicyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeletePolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeletePolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeletePolicyNoContent creates a DeletePolicyNoContent with default headers values
func NewDeletePolicyNoContent() *DeletePolicyNoContent {
	return &DeletePolicyNoContent{}
}

/*DeletePolicyNoContent handles this case with default header values.

Policy deleted
*/
type DeletePolicyNoContent struct {
}

func (o *DeletePolicyNoContent) Error() string {
	return fmt.Sprintf("[DELETE /access_policies/{id}][%d] deletePolicyNoContent ", 204)
}

func (o *DeletePolicyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePolicyUnauthorized creates a DeletePolicyUnauthorized with default headers values
func NewDeletePolicyUnauthorized() *DeletePolicyUnauthorized {
	return &DeletePolicyUnauthorized{}
}

/*DeletePolicyUnauthorized handles this case with default header values.

unauthorized: invalid credentials or missing authentication headers
*/
type DeletePolicyUnauthorized struct {
	Payload *models.GenericUnauthorizedResponse
}

func (o *DeletePolicyUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /access_policies/{id}][%d] deletePolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *DeletePolicyUnauthorized) GetPayload() *models.GenericUnauthorizedResponse {
	return o.Payload
}

func (o *DeletePolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericUnauthorizedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePolicyNotFound creates a DeletePolicyNotFound with default headers values
func NewDeletePolicyNotFound() *DeletePolicyNotFound {
	return &DeletePolicyNotFound{}
}

/*DeletePolicyNotFound handles this case with default header values.

policy not found
*/
type DeletePolicyNotFound struct {
}

func (o *DeletePolicyNotFound) Error() string {
	return fmt.Sprintf("[DELETE /access_policies/{id}][%d] deletePolicyNotFound ", 404)
}

func (o *DeletePolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
