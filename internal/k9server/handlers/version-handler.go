package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/kube9/k9/pkg/gen/k9server/server/operations"
	"github.com/kube9/k9/pkg/gen/k9server/server/operations/info"
	"github.com/kube9/k9/pkg/version"
)

// VersionHandler ...
type VersionHandler struct {
}

// NewVersionHandler ctor function
func NewVersionHandler() *VersionHandler {
	return &VersionHandler{}
}

// RegisterEndpoints ...
func (z *VersionHandler) RegisterEndpoints(openapi *operations.K9ServerAPI) {
	openapi.InfoGetVersionHandler = info.GetVersionHandlerFunc(z.versionHandler)
}

func (z *VersionHandler) versionHandler(params info.GetVersionParams) middleware.Responder {
	return info.NewGetVersionOK().WithPayload(version.Version)
}
