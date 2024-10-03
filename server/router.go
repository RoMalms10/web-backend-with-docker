package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"Simple-Web-Backend-With-DB-ChatGPT/handler"
)

func NewRouter(db *gorm.DB) *mux.Router {
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/people", handler.GetPeopleHandler(db)).Methods(http.MethodGet)
	r.HandleFunc("/people", handler.PostPeopleHandler(db)).Methods(http.MethodPost)
	r.HandleFunc("/people/{name}", handler.DeletePersonHandler(db)).Methods(http.MethodDelete)

	return r
}
