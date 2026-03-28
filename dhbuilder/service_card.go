package dhbuilder

import (
	"context"

	"github.com/farkaz00/dh_builder_service/dhbuilder/models"
)

func (srv *DHService) CreateCard(ctx context.Context, card *models.Card) (string, error) {
	return "", nil
}

func (srv *DHService) UpdateCard(ctx context.Context, card *models.Card) (string, error) {
	return "", nil
}

func (srv *DHService) GetCard(ctx context.Context, cardID string) (*models.Card, error) {
	return nil, nil
}

func (srv *DHService) GetCards(ctx context.Context) ([]*models.Card, error) {
	return nil, nil
}
