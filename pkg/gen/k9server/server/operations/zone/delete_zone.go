// Code generated by go-swagger; DO NOT EDIT.

package zone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteZoneHandlerFunc turns a function with the right signature into a delete zone handler
type DeleteZoneHandlerFunc func(DeleteZoneParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteZoneHandlerFunc) Handle(params DeleteZoneParams) middleware.Responder {
	return fn(params)
}

// DeleteZoneHandler interface for that can handle valid delete zone params
type DeleteZoneHandler interface {
	Handle(DeleteZoneParams) middleware.Responder
}

// NewDeleteZone creates a new http.Handler for the delete zone operation
func NewDeleteZone(ctx *middleware.Context, handler DeleteZoneHandler) *DeleteZone {
	return &DeleteZone{Context: ctx, Handler: handler}
}

/*DeleteZone swagger:route DELETE /zones/{zoneId} zone deleteZone

Delete a zone from k9 server

Delete a zone

*/
type DeleteZone struct {
	Context *middleware.Context
	Handler DeleteZoneHandler
}

func (o *DeleteZone) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteZoneParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
