package main

import (
	"log"
	"net/http"

	dhb "github.com/farkaz00/dh_builder_service/dhbuilder"
	dhdao "github.com/farkaz00/dh_builder_service/dhbuilder_dao"
	"github.com/farkaz00/dh_builder_service/dhbuilder_server/rest"

	"go.uber.org/zap"
)

func main() {
	// TODO: Add procedure to parse input flags as parameters

	// Create logger
	// TODO: Define logger level parameters
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal("failed to create service logger")
	}

	// Create DAO
	dhDAO, err := dhdao.NewDHDAO(dhdao.DAOTypeCSV)
	if err != nil {
		log.Fatal("failed to create service dao")
	}

	// Bundle service dependencies
	deps := &dhb.DHServiceDeps{
		Logger: logger,
		DAO:    dhDAO,
	}

	// Create service instance
	srv := dhb.NewDHService(deps)

	// Create router
	r := rest.NewRouter(srv)

	// Start http interface
	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
