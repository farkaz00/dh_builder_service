package rest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/farkaz00/dh_builder_service/dhbuilder/models"
	"github.com/gorilla/mux"
)

// IDRequest is a generic request model for endpoints that require only a resource ID
type IDRequest struct {
	ID string
}

func decodeIDRequest(ctx context.Context, r *http.Request) (any, error) {
	return IDRequest{ID: mux.Vars(r)["id"]}, nil
}

func decodeCardRequest(ctx context.Context, r *http.Request) (any, error) {
	var card models.Card
	if err := json.NewDecoder(r.Body).Decode(&card); err != nil {
		return nil, err
	}
	return &card, nil
}

func decodeUpdateCardRequest(ctx context.Context, r *http.Request) (any, error) {
	var card models.Card
	if err := json.NewDecoder(r.Body).Decode(&card); err != nil {
		return nil, err
	}
	card.ID = mux.Vars(r)["id"]
	return &card, nil
}

func decodeNoRequest(ctx context.Context, r *http.Request) (any, error) {
	return nil, nil
}
