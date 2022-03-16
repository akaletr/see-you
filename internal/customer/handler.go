package customer

import (
	"net/http"

	"cmd/main/main.go/internal/handlers"

	"github.com/go-chi/chi/v5"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (handler *handler) Register(router *chi.Mux) {
	router.Get("/customers", handler.GetAllCustomers)
	router.Post("/customers", handler.CreateCustomer)
}

func (handler *handler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {

}

func (handler *handler) CreateCustomer(w http.ResponseWriter, r *http.Request) {

}
