package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"Simple-Web-Backend-With-DB-ChatGPT/models"
)

func GetPeopleHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name != "" {
			// Get a specific person by name
			var p models.Person
			result := db.Where("name = ?", name).First(&p)
			if result.Error != nil {
				http.Error(w, "Person not found.", http.StatusNotFound)
				return
			}
			json.NewEncoder(w).Encode(p)
			return
		}

		// Return the entire list of people
		var people []models.Person
		db.Find(&people)
		json.NewEncoder(w).Encode(people)
	}
}

func PostPeopleHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var person models.Person
		err := json.NewDecoder(r.Body).Decode(&person)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		db.Create(&person)
		w.WriteHeader(http.StatusCreated)
	}
}

func DeletePersonHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]

		result := db.Delete(&models.Person{}, "name = ?", name)
		if result.RowsAffected == 0 {
			http.Error(w, "Person not found.", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
