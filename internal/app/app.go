package app

import (
	"net/http"

	"cmd/main/main.go/internal/config"
	"cmd/main/main.go/internal/storage"

	"github.com/go-chi/chi/v5"
)

type app struct {
	storage storage.Storage
	cfg     *config.Config
}

func New(config *config.Config) App {
	return &app{
		storage: storage.New(),
		cfg:     config,
	}
}

func (app *app) Start() error {
	router := chi.NewRouter()

	router.Post("/api/v1.1", app.Handle)

	server := http.Server{
		Addr:    app.cfg.ServerAddress,
		Handler: router,
	}

	return server.ListenAndServe()
}

func (app *app) Handle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
