package csv

import (
	"context"
	"os"
	"testing"

	"github.com/farkaz00/dh_builder_service/dhbuilder/models"
)

func newTestDAO(t *testing.T) (*DHCSV, func()) {
	t.Helper()
	f, err := os.CreateTemp("", "dh_cards_*.csv")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	path := f.Name()
	f.Close()
	os.Remove(path) // start with no file so DAO initialises from scratch
	return &DHCSV{filePath: path}, func() { os.Remove(path) }
}

var testCard = &models.Card{
	ID:           "card-1",
	Name:         "Shadow Strike",
	ManaCost:     3,
	Effect:       "Deal 2 damage",
	Image:        "shadow_strike.png",
	Realm:        models.DistortedShadows,
	LimitPerDeck: 2,
}

// SaveCard

func TestSaveCard_GeneratesIDWhenEmpty(t *testing.T) {
	dao, cleanup := newTestDAO(t)
	defer cleanup()

	card := &models.Card{Name: "Shadow Strike", ManaCost: 3}
	id, err := dao.SaveCard(context.Background(), card)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if id == "" {
		t.Error("expected a generated ID, got empty string")
	}
	if card.ID != id {
		t.Errorf("expected card.ID to be updated to %q, got %q", id, card.ID)
	}
}

func TestSaveCard_PreservesProvidedID(t *testing.T) {
	dao, cleanup := newTestDAO(t)
	defer cleanup()

	id, err := dao.SaveCard(context.Background(), testCard)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if id != testCard.ID {
		t.Errorf("expected id %q, got %q", testCard.ID, id)
	}
}

func TestSaveCard_UpdatesExistingCard(t *testing.T) {
	dao, cleanup := newTestDAO(t)
	defer cleanup()

	if _, err := dao.SaveCard(context.Background(), testCard); err != nil {
		t.Fatalf("setup SaveCard failed: %v", err)
	}

	updated := &models.Card{
		ID:           testCard.ID,
		Name:         "Shadow Strike II",
		ManaCost:     5,
		Effect:       "Deal 4 damage",
		Image:        testCard.Image,
		Realm:        testCard.Realm,
		LimitPerDeck: testCard.LimitPerDeck,
	}
	if _, err := dao.SaveCard(context.Background(), updated); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	cards, err := dao.GetCards(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(cards) != 1 {
		t.Fatalf("expected 1 card, got %d", len(cards))
	}
	if cards[0].Name != updated.Name {
		t.Errorf("expected name %q, got %q", updated.Name, cards[0].Name)
	}
	if cards[0].ManaCost != updated.ManaCost {
		t.Errorf("expected mana_cost %d, got %d", updated.ManaCost, cards[0].ManaCost)
	}
}

// GetCard

func TestGetCard_ReturnsCardByID(t *testing.T) {
	dao, cleanup := newTestDAO(t)
	defer cleanup()

	if _, err := dao.SaveCard(context.Background(), testCard); err != nil {
		t.Fatalf("setup SaveCard failed: %v", err)
	}

	card, err := dao.GetCard(context.Background(), testCard.ID)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if card == nil {
		t.Fatal("expected card, got nil")
	}
	if card.ID != testCard.ID {
		t.Errorf("expected id %q, got %q", testCard.ID, card.ID)
	}
	if card.Name != testCard.Name {
		t.Errorf("expected name %q, got %q", testCard.Name, card.Name)
	}
}

func TestGetCard_ReturnsNilWhenNotFound(t *testing.T) {
	dao, cleanup := newTestDAO(t)
	defer cleanup()

	if _, err := dao.SaveCard(context.Background(), testCard); err != nil {
		t.Fatalf("setup SaveCard failed: %v", err)
	}

	card, err := dao.GetCard(context.Background(), "unknown-id")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if card != nil {
		t.Errorf("expected nil, got %v", card)
	}
}

func TestGetCard_ReturnsNilWhenFileNotExists(t *testing.T) {
	dao, cleanup := newTestDAO(t)
	defer cleanup()

	card, err := dao.GetCard(context.Background(), "any-id")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if card != nil {
		t.Errorf("expected nil, got %v", card)
	}
}

// GetCards

func TestGetCards_ReturnsAllSavedCards(t *testing.T) {
	dao, cleanup := newTestDAO(t)
	defer cleanup()

	second := &models.Card{ID: "card-2", Name: "Wind Slash", ManaCost: 2, Realm: models.PiercingWinds, LimitPerDeck: 3}
	for _, c := range []*models.Card{testCard, second} {
		if _, err := dao.SaveCard(context.Background(), c); err != nil {
			t.Fatalf("setup SaveCard failed: %v", err)
		}
	}

	cards, err := dao.GetCards(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(cards) != 2 {
		t.Errorf("expected 2 cards, got %d", len(cards))
	}
}

func TestGetCards_ReturnsEmptyWhenFileNotExists(t *testing.T) {
	dao, cleanup := newTestDAO(t)
	defer cleanup()

	cards, err := dao.GetCards(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(cards) != 0 {
		t.Errorf("expected 0 cards, got %d", len(cards))
	}
}
