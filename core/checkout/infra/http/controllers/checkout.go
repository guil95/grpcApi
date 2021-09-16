package http

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/guil95/grpcApi/core/checkout/domain"
	"log"
	"net/http"
)

func MakeCheckoutHandler(r *mux.Router, service domain.CreateCheckoutUseCase) {
	r.Handle("/checkout", checkout(service)).Methods("POST", "OPTIONS").Name("create_checkout")
}

func checkout(service domain.CreateCheckoutUseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		var chartPayload domain.Chart

		if err := json.NewDecoder(r.Body).Decode(&chartPayload); err != nil {
			log.Println(err.Error())
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(newResponseError("Internal server error"))
			return
		}

		if !isValidRequest(chartPayload) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(newResponseError("Unprocessable entity"))
			return
		}

		err := chartPayload.Validate()

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(newResponseError("Unprocessable entity"))
			return
		}

		order, err := service.Checkout(&chartPayload)

		if err != nil {
			log.Println(err.Error())

			switch err {
			case domain.ProductGiftError:
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(newResponseError("Have a product gift in chart"))

				return
			case domain.ProductNotFoundError :
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(newResponseError("Products not found"))

				return
			default:
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(newResponseError("Internal server error"))

				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(order)

		return
	})
}

type ResponseError struct {
	Message string `json:"message"`
}

func newResponseError(message string) ResponseError {
	return ResponseError{
		Message: message,
	}
}

func isValidRequest(payload interface{}) bool {
	govalidator.SetFieldsRequiredByDefault(true)
	if _, err := govalidator.ValidateStruct(payload); err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}