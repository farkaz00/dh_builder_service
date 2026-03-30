package rest

import (
	"context"
	"net/http"

	dhb "github.com/farkaz00/dh_builder_service/dhbuilder"
	"github.com/farkaz00/dh_builder_service/dhbuilder/models"
)

type (
	// Handler maker inputs
	serviceCallerFunction      func(ctx context.Context, data any) (any, error)
	serviceCallerMakerFunction func(srv dhb.DHServicer) serviceCallerFunction
	requestDecoderFunction     func(ctx context.Context, r *http.Request) (any, error)
	requestEncoderFunction     func(ctx context.Context, data any, w http.ResponseWriter) (any, error)
	errorHandlingFunction      func(ctx context.Context, err error, w http.ResponseWriter) (any, error)
)

// HandlerWrapper Streamlines the handler inner data flow.
// Forces the use of request decoder and response encoder functions.
// Streamlines the overall error handling.
func HandlerWrapper(
	srv dhb.DHServicer,
	serviceCallerMakerFunc serviceCallerMakerFunction,
	requestDecoderFunc requestDecoderFunction,
	responseEncoderFunc requestEncoderFunction,
	errorHandlingFunc errorHandlingFunction,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Build Context and bundle the user session
		ctx := context.Background()

		data, err := requestDecoderFunc(ctx, r)
		if err != nil {
			errorHandlingFunc(ctx, err, w)
			return
		}

		toEncode, err := serviceCallerMakerFunc(srv)(ctx, data)
		if err != nil {
			errorHandlingFunc(ctx, err, w)
			return
		}
		_, err = responseEncoderFunc(ctx, toEncode, w)
		if err != nil {
			errorHandlingFunc(ctx, err, w)
			return
		}
	}
}

func CreateCardHandler(srv dhb.DHServicer) serviceCallerFunction {
	return func(ctx context.Context, data any) (any, error) {
		return srv.CreateCard(ctx, data.(*models.Card))
	}
}

func UpdateCardHandler(srv dhb.DHServicer) serviceCallerFunction {
	return func(ctx context.Context, data any) (any, error) {
		return srv.UpdateCard(ctx, data.(*models.Card))
	}
}

func GetCardHandler(srv dhb.DHServicer) serviceCallerFunction {
	return func(ctx context.Context, data any) (any, error) {
		return srv.GetCard(ctx, data.(IDRequest).ID)
	}
}

func GetCardsHandler(srv dhb.DHServicer) serviceCallerFunction {
	return func(ctx context.Context, data any) (any, error) {
		return srv.GetCards(ctx)
	}
}
