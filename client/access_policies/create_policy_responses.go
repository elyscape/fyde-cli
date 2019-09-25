// Code generated by go-swagger; DO NOT EDIT.

package access_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/oNaiPs/fyde-cli/models"
)

// CreatePolicyReader is a Reader for the CreatePolicy structure.
type CreatePolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreatePolicyCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreatePolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreatePolicyCreated creates a CreatePolicyCreated with default headers values
func NewCreatePolicyCreated() *CreatePolicyCreated {
	return &CreatePolicyCreated{}
}

/*CreatePolicyCreated handles this case with default header values.

Policy created
*/
type CreatePolicyCreated struct {
	Payload *CreatePolicyCreatedBody
}

func (o *CreatePolicyCreated) Error() string {
	return fmt.Sprintf("[POST /access_policies][%d] createPolicyCreated  %+v", 201, o.Payload)
}

func (o *CreatePolicyCreated) GetPayload() *CreatePolicyCreatedBody {
	return o.Payload
}

func (o *CreatePolicyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreatePolicyCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePolicyUnauthorized creates a CreatePolicyUnauthorized with default headers values
func NewCreatePolicyUnauthorized() *CreatePolicyUnauthorized {
	return &CreatePolicyUnauthorized{}
}

/*CreatePolicyUnauthorized handles this case with default header values.

unauthorized: invalid credentials or missing authentication headers
*/
type CreatePolicyUnauthorized struct {
	Payload *models.GenericUnauthorizedResponse
}

func (o *CreatePolicyUnauthorized) Error() string {
	return fmt.Sprintf("[POST /access_policies][%d] createPolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *CreatePolicyUnauthorized) GetPayload() *models.GenericUnauthorizedResponse {
	return o.Payload
}

func (o *CreatePolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericUnauthorizedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AccessResourcesItems0 access resources items0
swagger:model AccessResourcesItems0
*/
type AccessResourcesItems0 struct {

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this access resources items0
func (o *AccessResourcesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AccessResourcesItems0) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(o.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", o.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *AccessResourcesItems0) validateID(formats strfmt.Registry) error {

	if swag.IsZero(o.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", o.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *AccessResourcesItems0) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(o.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", o.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AccessResourcesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AccessResourcesItems0) UnmarshalBinary(b []byte) error {
	var res AccessResourcesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreatePolicyBody create policy body
swagger:model CreatePolicyBody
*/
type CreatePolicyBody struct {

	// access policy
	AccessPolicy *CreatePolicyParamsBodyAccessPolicy `json:"access_policy,omitempty"`
}

// Validate validates this create policy body
func (o *CreatePolicyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAccessPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreatePolicyBody) validateAccessPolicy(formats strfmt.Registry) error {

	if swag.IsZero(o.AccessPolicy) { // not required
		return nil
	}

	if o.AccessPolicy != nil {
		if err := o.AccessPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy" + "." + "access_policy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreatePolicyBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreatePolicyBody) UnmarshalBinary(b []byte) error {
	var res CreatePolicyBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreatePolicyCreatedBody create policy created body
swagger:model CreatePolicyCreatedBody
*/
type CreatePolicyCreatedBody struct {
	models.AccessPolicy

	// access resources
	AccessResources []*AccessResourcesItems0 `json:"access_resources"`

	// groups
	Groups []*GroupsItems0 `json:"groups"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *CreatePolicyCreatedBody) UnmarshalJSON(raw []byte) error {
	// CreatePolicyCreatedBodyAO0
	var createPolicyCreatedBodyAO0 models.AccessPolicy
	if err := swag.ReadJSON(raw, &createPolicyCreatedBodyAO0); err != nil {
		return err
	}
	o.AccessPolicy = createPolicyCreatedBodyAO0

	// CreatePolicyCreatedBodyAO1
	var dataCreatePolicyCreatedBodyAO1 struct {
		AccessResources []*AccessResourcesItems0 `json:"access_resources"`

		Groups []*GroupsItems0 `json:"groups"`
	}
	if err := swag.ReadJSON(raw, &dataCreatePolicyCreatedBodyAO1); err != nil {
		return err
	}

	o.AccessResources = dataCreatePolicyCreatedBodyAO1.AccessResources

	o.Groups = dataCreatePolicyCreatedBodyAO1.Groups

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o CreatePolicyCreatedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	createPolicyCreatedBodyAO0, err := swag.WriteJSON(o.AccessPolicy)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, createPolicyCreatedBodyAO0)

	var dataCreatePolicyCreatedBodyAO1 struct {
		AccessResources []*AccessResourcesItems0 `json:"access_resources"`

		Groups []*GroupsItems0 `json:"groups"`
	}

	dataCreatePolicyCreatedBodyAO1.AccessResources = o.AccessResources

	dataCreatePolicyCreatedBodyAO1.Groups = o.Groups

	jsonDataCreatePolicyCreatedBodyAO1, errCreatePolicyCreatedBodyAO1 := swag.WriteJSON(dataCreatePolicyCreatedBodyAO1)
	if errCreatePolicyCreatedBodyAO1 != nil {
		return nil, errCreatePolicyCreatedBodyAO1
	}
	_parts = append(_parts, jsonDataCreatePolicyCreatedBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this create policy created body
func (o *CreatePolicyCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.AccessPolicy
	if err := o.AccessPolicy.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateAccessResources(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateGroups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreatePolicyCreatedBody) validateAccessResources(formats strfmt.Registry) error {

	if swag.IsZero(o.AccessResources) { // not required
		return nil
	}

	for i := 0; i < len(o.AccessResources); i++ {
		if swag.IsZero(o.AccessResources[i]) { // not required
			continue
		}

		if o.AccessResources[i] != nil {
			if err := o.AccessResources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createPolicyCreated" + "." + "access_resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *CreatePolicyCreatedBody) validateGroups(formats strfmt.Registry) error {

	if swag.IsZero(o.Groups) { // not required
		return nil
	}

	for i := 0; i < len(o.Groups); i++ {
		if swag.IsZero(o.Groups[i]) { // not required
			continue
		}

		if o.Groups[i] != nil {
			if err := o.Groups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createPolicyCreated" + "." + "groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreatePolicyCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreatePolicyCreatedBody) UnmarshalBinary(b []byte) error {
	var res CreatePolicyCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreatePolicyParamsBodyAccessPolicy create policy params body access policy
swagger:model CreatePolicyParamsBodyAccessPolicy
*/
type CreatePolicyParamsBodyAccessPolicy struct {

	// access resource ids
	AccessResourceIds []strfmt.UUID `json:"access_resource_ids"`

	// conditions
	Conditions interface{} `json:"conditions,omitempty"`

	// group ids
	GroupIds []int64 `json:"group_ids"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this create policy params body access policy
func (o *CreatePolicyParamsBodyAccessPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAccessResourceIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreatePolicyParamsBodyAccessPolicy) validateAccessResourceIds(formats strfmt.Registry) error {

	if swag.IsZero(o.AccessResourceIds) { // not required
		return nil
	}

	for i := 0; i < len(o.AccessResourceIds); i++ {

		if err := validate.FormatOf("policy"+"."+"access_policy"+"."+"access_resource_ids"+"."+strconv.Itoa(i), "body", "uuid", o.AccessResourceIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreatePolicyParamsBodyAccessPolicy) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreatePolicyParamsBodyAccessPolicy) UnmarshalBinary(b []byte) error {
	var res CreatePolicyParamsBodyAccessPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GroupsItems0 groups items0
swagger:model GroupsItems0
*/
type GroupsItems0 struct {

	// color
	Color *string `json:"color,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this groups items0
func (o *GroupsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GroupsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GroupsItems0) UnmarshalBinary(b []byte) error {
	var res GroupsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
