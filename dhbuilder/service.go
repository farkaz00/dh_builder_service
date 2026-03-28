// Package dhbuilder service layer, providing business logic and rule validations
package dhbuilder

import (
	"fmt"

	"go.uber.org/zap"
)

// DHServiceDeps bundles the dependencies to build the service instance
type DHServiceDeps struct {
	Logger *zap.Logger
	DAO    DHDAO
}

// DHService stores all servcie dependencies
type DHService struct {
	logger *zap.Logger
	dao    DHDAO
}

func NewDHService(deps *DHServiceDeps) DHServicer {
	return &DHService{
		logger: deps.Logger,
		dao:    deps.DAO,
	}
}

func (dhs *DHService) Serve() {
	fmt.Println("Serving...")
}
