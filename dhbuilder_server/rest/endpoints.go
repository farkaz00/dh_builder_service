package rest

import (
	"net/http"

	dhb "github.com/farkaz00/dh_builder_service/dhbuilder"
	"github.com/gorilla/mux"
)

// NewRouter builds all the routes used by the rest interface
func NewRouter(srv dhb.DHServicer) *mux.Router {
	// Create Handler functions
	loginHandler := HandlerWrapper(
		srv,
		CreateCardHandler,
		decodeLoginRequest,
		JSONEncoder,
		ServerErrorEncoder,
	)

	// Assign Handler functions
	r := mux.NewRouter()
	r.HandleFunc("/api/dhbuilder/cards", loginHandler).Methods(http.MethodPost)
	return r
}
