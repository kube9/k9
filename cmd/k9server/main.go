package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/kube9/k9/internal/k9server/apis"
	"github.com/kube9/k9/internal/k9server/datas"
	"github.com/kube9/k9/internal/k9server/handlers"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-openapi/loads"
	"github.com/kube9/k9/pkg/gen/k9server/server"
	"github.com/kube9/k9/pkg/gen/k9server/server/operations"
)

func main() {
	// setup prometheus
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		panic(http.ListenAndServe(":9090", nil))
	}()

	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(server.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// create new service API
	openapi := operations.NewK9ServerAPI(swaggerSpec)
	server := server.NewServer(openapi)
	defer server.Shutdown()

	// parse flags
	flag.Parse()
	// set the port this service will be run on
	server.Port = 3000

	// Create stores
	zoneData := datas.NewZoneDataMock()

	// Create apis
	zoneAPI := apis.NewZoneAPI(zoneData)

	// Create handlers
	versionHandler := handlers.NewVersionHandler()
	versionHandler.RegisterEndpoints(openapi)
	zoneHandler := handlers.NewZoneHandler(zoneAPI)
	zoneHandler.RegisterEndpoints(openapi)

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
