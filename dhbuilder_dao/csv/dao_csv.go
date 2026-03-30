// Package csv implements the DHDAO interface using CSV files for persistence
package csv

import (
	dhb "github.com/farkaz00/dh_builder_service/dhbuilder"
)

// DHCSVDeps bundles the dependencies to build the CSV DAO instance
type DHCSVDeps struct {
	FilePath string
}

// DHCSV stores all CSV DAO dependencies
type DHCSV struct {
	filePath string
}

func NewDHCSV(deps *DHCSVDeps) (dhb.DHDAO, error) {
	return &DHCSV{
		filePath: deps.FilePath,
	}, nil
}
