package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func JSONEncoder(ctx context.Context, data any, w http.ResponseWriter) (any, error) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	return nil, err
}

func ServerErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) (any, error) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusInternalServerError)
	_, writeErr := fmt.Fprint(w, err.Error())
	return nil, writeErr
}
