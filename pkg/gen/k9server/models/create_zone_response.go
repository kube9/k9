// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CreateZoneResponse create zone response
// swagger:model CreateZoneResponse
type CreateZoneResponse struct {
	Zone
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *CreateZoneResponse) UnmarshalJSON(raw []byte) error {

	var aO0 Zone
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Zone = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m CreateZoneResponse) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.Zone)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this create zone response
func (m *CreateZoneResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.Zone.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CreateZoneResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateZoneResponse) UnmarshalBinary(b []byte) error {
	var res CreateZoneResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
