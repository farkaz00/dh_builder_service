package rest

import (
	"net/http"

	dhb "github.com/farkaz00/dh_builder_service/dhbuilder"
	"github.com/gorilla/mux"
)

// NewRouter builds all the routes used by the rest interface
func NewRouter(srv dhb.DHServicer) *mux.Router {
	// Create Handler functions
	cardCreateHandler := HandlerWrapper(srv, CreateCardHandler, decodeCardRequest, JSONEncoder, ServerErrorEncoder)
	cardUpdateHandler := HandlerWrapper(srv, UpdateCardHandler, decodeUpdateCardRequest, JSONEncoder, ServerErrorEncoder)
	cardGetHandler := HandlerWrapper(srv, GetCardHandler, decodeIDRequest, JSONEncoder, ServerErrorEncoder)
	cardGetAllHandler := HandlerWrapper(srv, GetCardsHandler, decodeNoRequest, JSONEncoder, ServerErrorEncoder)

	// Assign Handler functions
	r := mux.NewRouter()

	// Cards
	r.HandleFunc("/api/dhbuilder/cards", cardCreateHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/dhbuilder/cards/{id}", cardUpdateHandler).Methods(http.MethodPut)
	r.HandleFunc("/api/dhbuilder/cards/{id}", cardGetHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/dhbuilder/cards", cardGetAllHandler).Methods(http.MethodGet)
	return r
}
