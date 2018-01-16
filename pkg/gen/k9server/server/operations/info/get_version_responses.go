// Code generated by go-swagger; DO NOT EDIT.

package info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetVersionOKCode is the HTTP code returned for type GetVersionOK
const GetVersionOKCode int = 200

/*GetVersionOK ok

swagger:response getVersionOK
*/
type GetVersionOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetVersionOK creates GetVersionOK with default headers values
func NewGetVersionOK() *GetVersionOK {
	return &GetVersionOK{}
}

// WithPayload adds the payload to the get version o k response
func (o *GetVersionOK) WithPayload(payload string) *GetVersionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get version o k response
func (o *GetVersionOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVersionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}