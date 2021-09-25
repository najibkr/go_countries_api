package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/najibkr/go_countries_api/entities"
)

var (
	countries []entities.Country
)

func init() {
	countries = []entities.Country{
		{ID: "1", Name: "Afghanistan", Currency: entities.Currency{Code: "AFN", Name: "Afghani", Symbol: "AFN"}},
		{ID: "1", Name: "Afghanistan", Currency: entities.Currency{Code: "AFN", Name: "Afghani", Symbol: "AFN"}},
		{ID: "1", Name: "Afghanistan", Currency: entities.Currency{Code: "AFN", Name: "Afghani", Symbol: "AFN"}},
		{ID: "1", Name: "Afghanistan", Currency: entities.Currency{Code: "AFN", Name: "Afghani", Symbol: "AFN"}},
		{ID: "1", Name: "Afghanistan", Currency: entities.Currency{Code: "AFN", Name: "Afghani", Symbol: "AFN"}},
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", countriesHandler).Methods("GET")
	fmt.Println("Listening and Serving on: http://localhost:8080")
	http.ListenAndServe(":8080", router)
}

func countriesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, err := json.Marshal(countries)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"could not encode json"`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}
