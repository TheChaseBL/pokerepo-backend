package api

import (
	"encoding/json"
	"github.com/TheChaseBL/pokerepo-backend/database"
	"github.com/TheChaseBL/pokerepo-backend/entity"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func AddPokemonHandler(w http.ResponseWriter, r *http.Request) {
	var p entity.Pokemon
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	database.Conn.AddEntry(p)
}

func ReadPokemonHandler(w http.ResponseWriter, _ *http.Request) {
	data, err := json.Marshal(database.Conn.ReadAll())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func StartAPI() {
	r := mux.NewRouter()
	r.HandleFunc("/api/pokemon/add", AddPokemonHandler).Methods("POST")
	r.HandleFunc("/api/pokemon", ReadPokemonHandler).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}