package dhbuilder

import (
	"context"
	"errors"

	"github.com/farkaz00/dh_builder_service/dhbuilder/models"
	"go.uber.org/zap"
)

var (
	ErrCardAlreadyExists = errors.New("card already exists")
	ErrCardNotFound      = errors.New("card not found")
)

func (dhs *DHService) CreateCard(ctx context.Context, card *models.Card) (string, error) {
	logger := dhs.logger.With(zap.String("method", "CreateCard"))

	if card.ID != "" {
		existing, err := dhs.dao.GetCard(ctx, card.ID)
		if err != nil {
			logger.Error("failed to check card existence", zap.Error(err))
			return "", err
		}
		if existing != nil {
			logger.Error("card already exists", zap.String("card_id", card.ID))
			return "", ErrCardAlreadyExists
		}
	}

	id, err := dhs.dao.SaveCard(ctx, card)
	if err != nil {
		logger.Error("failed to save card", zap.Error(err))
		return "", err
	}
	return id, nil
}

func (dhs *DHService) UpdateCard(ctx context.Context, card *models.Card) (string, error) {
	logger := dhs.logger.With(zap.String("method", "UpdateCard"))

	existing, err := dhs.dao.GetCard(ctx, card.ID)
	if err != nil {
		logger.Error("failed to check card existence", zap.Error(err))
		return "", err
	}
	if existing == nil {
		logger.Error("card not found", zap.String("card_id", card.ID))
		return "", ErrCardNotFound
	}

	id, err := dhs.dao.SaveCard(ctx, card)
	if err != nil {
		logger.Error("failed to save card", zap.Error(err))
		return "", err
	}
	return id, nil
}

func (dhs *DHService) GetCard(ctx context.Context, cardID string) (*models.Card, error) {
	logger := dhs.logger.With(zap.String("method", "GetCard"))

	card, err := dhs.dao.GetCard(ctx, cardID)
	if err != nil {
		logger.Error("failed to get card", zap.Error(err))
		return nil, err
	}
	return card, nil
}

func (dhs *DHService) GetCards(ctx context.Context) ([]*models.Card, error) {
	logger := dhs.logger.With(zap.String("method", "GetCards"))

	cards, err := dhs.dao.GetCards(ctx)
	if err != nil {
		logger.Error("failed to get cards", zap.Error(err))
		return nil, err
	}
	return cards, nil
}
