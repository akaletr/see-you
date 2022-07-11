package app

import (
	"cmd/main/main.go/internal/storage"
	"cmd/main/main.go/internal/user"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"log"
	"net/http"
	"strconv"
)

type app struct {
	storage storage.Storage
}

func New() App {
	return &app{
		storage: storage.New(),
	}
}

func (app *app) Start() error {
	router := chi.NewRouter()

	router.Post("/add", app.AddUser)
	router.Get("/{id}", app.GetUser)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	return server.ListenAndServe()
}

func (app *app) AddUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer func() {
		err = r.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	u := user.User{}
	err = json.Unmarshal(body, &u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	fmt.Println(u)
	id, err := u.Save()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	fmt.Println(u)

	idString := strconv.FormatUint(id, 10)

	err = app.storage.Write(idString, u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(idString))
	if err != nil {
		log.Println(err)
	}
}

func (app *app) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	u, err := app.storage.Read(id)

	uJSON, err := json.Marshal(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(uJSON)
	if err != nil {
		log.Println(err)
	}
}
