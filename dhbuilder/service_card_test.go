package dhbuilder

import (
	"context"
	"errors"
	"testing"

	"github.com/farkaz00/dh_builder_service/dhbuilder/models"
	"go.uber.org/zap"
)

// mockDAO is a test double for the DHDAO interface
type mockDAO struct {
	saveCard func(ctx context.Context, card *models.Card) (string, error)
	getCard  func(ctx context.Context, cardID string) (*models.Card, error)
	getCards func(ctx context.Context) ([]*models.Card, error)
}

func (m *mockDAO) SaveCard(ctx context.Context, card *models.Card) (string, error) {
	return m.saveCard(ctx, card)
}

func (m *mockDAO) GetCard(ctx context.Context, cardID string) (*models.Card, error) {
	return m.getCard(ctx, cardID)
}

func (m *mockDAO) GetCards(ctx context.Context) ([]*models.Card, error) {
	return m.getCards(ctx)
}

func newTestService(dao DHDAO) *DHService {
	return &DHService{logger: zap.NewNop(), dao: dao}
}

var testCard = &models.Card{
	ID:           "card-1",
	Name:         "Shadow Strike",
	ManaCost:     3,
	Effect:       "Deal 2 damage",
	Realm:        models.DistortedShadows,
	LimitPerDeck: 2,
}

// CreateCard

func TestCreateCard_WithoutID_DelegatesToDAO(t *testing.T) {
	dao := &mockDAO{
		saveCard: func(_ context.Context, card *models.Card) (string, error) {
			return "generated-id", nil
		},
	}
	svc := newTestService(dao)

	id, err := svc.CreateCard(context.Background(), &models.Card{Name: "Shadow Strike"})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if id != "generated-id" {
		t.Errorf("expected id %q, got %q", "generated-id", id)
	}
}

func TestCreateCard_WithID_WhenNotExists_Saves(t *testing.T) {
	dao := &mockDAO{
		getCard: func(_ context.Context, _ string) (*models.Card, error) {
			return nil, nil
		},
		saveCard: func(_ context.Context, card *models.Card) (string, error) {
			return card.ID, nil
		},
	}
	svc := newTestService(dao)

	id, err := svc.CreateCard(context.Background(), testCard)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if id != testCard.ID {
		t.Errorf("expected id %q, got %q", testCard.ID, id)
	}
}

func TestCreateCard_WithID_WhenExists_ReturnsErrCardAlreadyExists(t *testing.T) {
	dao := &mockDAO{
		getCard: func(_ context.Context, _ string) (*models.Card, error) {
			return testCard, nil
		},
	}
	svc := newTestService(dao)

	_, err := svc.CreateCard(context.Background(), testCard)

	if !errors.Is(err, ErrCardAlreadyExists) {
		t.Errorf("expected ErrCardAlreadyExists, got %v", err)
	}
}

func TestCreateCard_GetCardError_ReturnsError(t *testing.T) {
	daoErr := errors.New("dao error")
	dao := &mockDAO{
		getCard: func(_ context.Context, _ string) (*models.Card, error) {
			return nil, daoErr
		},
	}
	svc := newTestService(dao)

	_, err := svc.CreateCard(context.Background(), testCard)

	if !errors.Is(err, daoErr) {
		t.Errorf("expected dao error, got %v", err)
	}
}

func TestCreateCard_SaveCardError_ReturnsError(t *testing.T) {
	daoErr := errors.New("save error")
	dao := &mockDAO{
		getCard: func(_ context.Context, _ string) (*models.Card, error) {
			return nil, nil
		},
		saveCard: func(_ context.Context, _ *models.Card) (string, error) {
			return "", daoErr
		},
	}
	svc := newTestService(dao)

	_, err := svc.CreateCard(context.Background(), testCard)

	if !errors.Is(err, daoErr) {
		t.Errorf("expected save error, got %v", err)
	}
}

// UpdateCard

func TestUpdateCard_WhenExists_Saves(t *testing.T) {
	dao := &mockDAO{
		getCard: func(_ context.Context, _ string) (*models.Card, error) {
			return testCard, nil
		},
		saveCard: func(_ context.Context, card *models.Card) (string, error) {
			return card.ID, nil
		},
	}
	svc := newTestService(dao)

	id, err := svc.UpdateCard(context.Background(), testCard)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if id != testCard.ID {
		t.Errorf("expected id %q, got %q", testCard.ID, id)
	}
}

func TestUpdateCard_WhenNotExists_ReturnsErrCardNotFound(t *testing.T) {
	dao := &mockDAO{
		getCard: func(_ context.Context, _ string) (*models.Card, error) {
			return nil, nil
		},
	}
	svc := newTestService(dao)

	_, err := svc.UpdateCard(context.Background(), testCard)

	if !errors.Is(err, ErrCardNotFound) {
		t.Errorf("expected ErrCardNotFound, got %v", err)
	}
}

func TestUpdateCard_GetCardError_ReturnsError(t *testing.T) {
	daoErr := errors.New("dao error")
	dao := &mockDAO{
		getCard: func(_ context.Context, _ string) (*models.Card, error) {
			return nil, daoErr
		},
	}
	svc := newTestService(dao)

	_, err := svc.UpdateCard(context.Background(), testCard)

	if !errors.Is(err, daoErr) {
		t.Errorf("expected dao error, got %v", err)
	}
}

func TestUpdateCard_SaveCardError_ReturnsError(t *testing.T) {
	daoErr := errors.New("save error")
	dao := &mockDAO{
		getCard: func(_ context.Context, _ string) (*models.Card, error) {
			return testCard, nil
		},
		saveCard: func(_ context.Context, _ *models.Card) (string, error) {
			return "", daoErr
		},
	}
	svc := newTestService(dao)

	_, err := svc.UpdateCard(context.Background(), testCard)

	if !errors.Is(err, daoErr) {
		t.Errorf("expected save error, got %v", err)
	}
}

// GetCard

func TestGetCard_ReturnsCard(t *testing.T) {
	dao := &mockDAO{
		getCard: func(_ context.Context, _ string) (*models.Card, error) {
			return testCard, nil
		},
	}
	svc := newTestService(dao)

	card, err := svc.GetCard(context.Background(), testCard.ID)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if card != testCard {
		t.Errorf("expected card %v, got %v", testCard, card)
	}
}

func TestGetCard_DAOError_ReturnsError(t *testing.T) {
	daoErr := errors.New("dao error")
	dao := &mockDAO{
		getCard: func(_ context.Context, _ string) (*models.Card, error) {
			return nil, daoErr
		},
	}
	svc := newTestService(dao)

	_, err := svc.GetCard(context.Background(), testCard.ID)

	if !errors.Is(err, daoErr) {
		t.Errorf("expected dao error, got %v", err)
	}
}

// GetCards

func TestGetCards_ReturnsCards(t *testing.T) {
	cards := []*models.Card{testCard}
	dao := &mockDAO{
		getCards: func(_ context.Context) ([]*models.Card, error) {
			return cards, nil
		},
	}
	svc := newTestService(dao)

	result, err := svc.GetCards(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != len(cards) {
		t.Errorf("expected %d cards, got %d", len(cards), len(result))
	}
}

func TestGetCards_DAOError_ReturnsError(t *testing.T) {
	daoErr := errors.New("dao error")
	dao := &mockDAO{
		getCards: func(_ context.Context) ([]*models.Card, error) {
			return nil, daoErr
		},
	}
	svc := newTestService(dao)

	_, err := svc.GetCards(context.Background())

	if !errors.Is(err, daoErr) {
		t.Errorf("expected dao error, got %v", err)
	}
}
