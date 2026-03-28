package main

import (
	"log"
	"net/http"

	dhb "github.com/farkaz00/dh_builder_service/dhbuilder"
	"github.com/farkaz00/dh_builder_service/dhbuilder_server/rest"

	"go.uber.org/zap"
)

func main() {
	// Create logger
	// TODO: Define logger level parameters
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal("failed to create service logger")
	}

	// Bundle service dependencies
	deps := &dhb.DHServiceDeps{
		Logger: logger,
	}

	// Create service instance
	srv := dhb.NewDHService(deps)

	// Create router
	r := rest.NewRouter(srv)

	// Start http interface
	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
