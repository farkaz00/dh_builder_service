package csv

import (
	"context"
	"crypto/rand"
	gocsv "encoding/csv"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"

	"github.com/farkaz00/dh_builder_service/dhbuilder/models"
)

var cardCSVHeaders = []string{"id", "mana_cost", "name", "description", "image", "card_realm", "limit_per_deck"}

func (dao *DHCSV) SaveCard(ctx context.Context, card *models.Card) (string, error) {
	if card.ID == "" {
		id, err := generateCardID()
		if err != nil {
			return "", err
		}
		card.ID = id
	}

	cards, err := dao.readAllCards()
	if err != nil {
		return "", err
	}

	found := false
	for i, c := range cards {
		if c.ID == card.ID {
			cards[i] = card
			found = true
			break
		}
	}
	if !found {
		cards = append(cards, card)
	}

	if err := dao.writeAllCards(cards); err != nil {
		return "", err
	}

	return card.ID, nil
}

func (dao *DHCSV) GetCard(ctx context.Context, cardID string) (*models.Card, error) {
	cards, err := dao.readAllCards()
	if err != nil {
		return nil, err
	}

	for _, card := range cards {
		if card.ID == cardID {
			return card, nil
		}
	}
	return nil, nil
}

func (dao *DHCSV) GetCards(ctx context.Context) ([]*models.Card, error) {
	return dao.readAllCards()
}

func (dao *DHCSV) readAllCards() ([]*models.Card, error) {
	f, err := os.Open(dao.cardFilePath)
	if os.IsNotExist(err) {
		return []*models.Card{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer f.Close()

	r := gocsv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV file: %w", err)
	}

	cards := make([]*models.Card, 0, len(records))
	for i, record := range records {
		if i == 0 { // skip header row
			continue
		}
		card, err := recordToCard(record)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}
	return cards, nil
}

func (dao *DHCSV) writeAllCards(cards []*models.Card) error {
	f, err := os.Create(dao.cardFilePath)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %w", err)
	}
	defer f.Close()

	w := gocsv.NewWriter(f)
	if err := w.Write(cardCSVHeaders); err != nil {
		return fmt.Errorf("failed to write CSV headers: %w", err)
	}
	for _, card := range cards {
		if err := w.Write(cardToRecord(card)); err != nil {
			return fmt.Errorf("failed to write card record: %w", err)
		}
	}
	w.Flush()
	return w.Error()
}

func recordToCard(record []string) (*models.Card, error) {
	manaCost, err := strconv.Atoi(record[1])
	if err != nil {
		return nil, fmt.Errorf("invalid mana_cost value %q: %w", record[1], err)
	}
	limitPerDeck, err := strconv.Atoi(record[6])
	if err != nil {
		return nil, fmt.Errorf("invalid limit_per_deck value %q: %w", record[6], err)
	}
	return &models.Card{
		ID:           record[0],
		ManaCost:     manaCost,
		Name:         record[2],
		Effect:       record[3],
		Image:        record[4],
		Realm:        models.CardRealm(record[5]),
		LimitPerDeck: limitPerDeck,
	}, nil
}

func cardToRecord(card *models.Card) []string {
	return []string{
		card.ID,
		strconv.Itoa(card.ManaCost),
		card.Name,
		card.Effect,
		card.Image,
		string(card.Realm),
		strconv.Itoa(card.LimitPerDeck),
	}
}

func generateCardID() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("failed to generate card ID: %w", err)
	}
	return hex.EncodeToString(b), nil
}
