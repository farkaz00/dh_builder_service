package dao

import (
	"fmt"

	dhb "github.com/farkaz00/dh_builder_service/dhbuilder"
	daocsv "github.com/farkaz00/dh_builder_service/dhbuilder_dao/csv"
)

// DAOType defines the supported Data Access Objects
type DAOType string

const (
	// Supported DAO types
	DAOTypeCSV DAOType = "csv"

	// Package errors
	DAOErrorUnsupportedDAOType = "unsupported DAO type"
)

func NewDHDAO(daoType DAOType) (dhb.DHDAO, error) {
	switch daoType {
	case DAOTypeCSV:
		return daocsv.NewDHCSV(&daocsv.DHCSVDeps{CardFilePath: "cards.csv"})
	default:
		return nil, fmt.Errorf("%s: %s", DAOErrorUnsupportedDAOType, daoType)
	}
}
