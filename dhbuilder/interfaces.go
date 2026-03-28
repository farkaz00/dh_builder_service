package dhbuilder

import (
	"context"

	"github.com/farkaz00/dh_builder_service/dhbuilder/models"
)

// DHServicer defines the minimum methods required to offer DH Builder capabilities
type DHServicer interface {
	// Session
	// TODO: Define session Management

	// Card
	CreateCard(ctx context.Context, card *models.Card) (string, error)
	UpdateCard(ctx context.Context, card *models.Card) (string, error)
	GetCard(ctx context.Context, cardID string) (*models.Card, error)
	GetCards(ctx context.Context) ([]*models.Card, error)
}

// DHDAO defines the required methods for Data Accesss Operations
type DHDAO interface {
	// Card
	SaveCard(ctx context.Context, card *models.Card) (string, error)
	GetCard(ctx context.Context, cardID string) (*models.Card, error)
	GetCards(ctx context.Context) ([]*models.Card, error)
}
