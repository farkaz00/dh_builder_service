package dhbuilder

import (
	"fmt"

	"go.uber.org/zap"
)

func (dhs *DHService) Login() error {
	logger := dhs.logger.With(
		zap.String("method", "Login"),
	)
	logger.Info("Login Started...")

	fmt.Print("Logged user")

	return nil
}
