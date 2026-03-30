package dhbuilder

import (
	"context"

	"github.com/farkaz00/dh_builder_service/dhbuilder/models"
)

func (dhs *DHService) CreateCard(ctx context.Context, card *models.Card) (string, error) {
	return "", nil
}

func (dhs *DHService) UpdateCard(ctx context.Context, card *models.Card) (string, error) {
	return "", nil
}

func (dhs *DHService) GetCard(ctx context.Context, cardID string) (*models.Card, error) {
	return nil, nil
}

func (dhs *DHService) GetCards(ctx context.Context) ([]*models.Card, error) {
	return nil, nil
}
