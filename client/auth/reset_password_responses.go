// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ResetPasswordReader is a Reader for the ResetPassword structure.
type ResetPasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResetPasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResetPasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewResetPasswordOK creates a ResetPasswordOK with default headers values
func NewResetPasswordOK() *ResetPasswordOK {
	return &ResetPasswordOK{}
}

/*ResetPasswordOK handles this case with default header values.

successful operation
*/
type ResetPasswordOK struct {
}

func (o *ResetPasswordOK) Error() string {
	return fmt.Sprintf("[POST /auth/password][%d] resetPasswordOK ", 200)
}

func (o *ResetPasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
