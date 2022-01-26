package place

import (
	"cmd/main/main.go/internal/handlers"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (handler *handler) Register(router *chi.Mux) {
	router.Get("/place", handler.GetAllPlaces)
	router.Post("/place", handler.CreatePlace)
}

func (handler *handler) GetAllPlaces(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(chi.URLParam(r, "path")))
	if err != nil {
		fmt.Println(err)
	}
}

func (handler *handler) CreatePlace(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("all places"))
}
