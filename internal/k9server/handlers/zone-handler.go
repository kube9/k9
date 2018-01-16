package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/kube9/k9/internal/k9server/apis"
	"github.com/kube9/k9/internal/k9server/errors"
	"github.com/kube9/k9/pkg/gen/k9server/models"
	"github.com/kube9/k9/pkg/gen/k9server/server/operations"
	"github.com/kube9/k9/pkg/gen/k9server/server/operations/zone"
)

// ZoneHandler ...
type ZoneHandler struct {
	api *apis.ZoneAPI
}

// NewZoneHandler ctor function
func NewZoneHandler(api *apis.ZoneAPI) *ZoneHandler {
	return &ZoneHandler{
		api,
	}
}

// RegisterEndpoints ...
func (z *ZoneHandler) RegisterEndpoints(openapi *operations.K9ServerAPI) {
	openapi.ZoneGetZonesHandler = zone.GetZonesHandlerFunc(z.getZonesHandler)
	openapi.ZoneCreateZoneHandler = zone.CreateZoneHandlerFunc(z.createZoneHandler)
	openapi.ZoneDeleteZoneHandler = zone.DeleteZoneHandlerFunc(z.deleteZoneHandler)
}

func (z *ZoneHandler) createZoneHandler(params zone.CreateZoneParams) middleware.Responder {
	data, err := z.api.CreateOne(params.Zone.NewZone)

	if err != nil {
		code := int32(500)
		msg := "Oeps..."
		return zone.NewCreateZoneInternalServerError().WithPayload(&models.ErrorResponse{
			ErrorCode:    &code,
			ErrorMessage: &msg,
		})
	}

	return zone.NewCreateZoneCreated().WithPayload(&models.CreateZoneResponse{
		Zone: *data,
	})
}

func (z *ZoneHandler) getZonesHandler(params zone.GetZonesParams) middleware.Responder {
	data, err := z.api.FindAll()

	if err != nil {
		code := int32(500)
		msg := "Oeps..."
		return zone.NewGetZonesInternalServerError().WithPayload(&models.ErrorResponse{
			ErrorCode:    &code,
			ErrorMessage: &msg,
		})
	}

	return zone.NewGetZonesOK().WithPayload(&models.GetZonesResponse{
		ZoneList: *data,
	})
}

func (z *ZoneHandler) deleteZoneHandler(params zone.DeleteZoneParams) middleware.Responder {
	err := z.api.DeleteOneByID(params.ZoneID)

	if err == errors.ErrIDNotFound {
		code := int32(404)
		msg := "ID Not Found"
		return zone.NewDeleteZoneNotFound().WithPayload(&models.ErrorResponse{
			ErrorCode:    &code,
			ErrorMessage: &msg,
		})
	}

	return zone.NewDeleteZoneNoContent()
}
