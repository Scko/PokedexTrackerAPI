package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Pokemon struct{
	ID		string		`json:"id,omitempty"`
	Name	string		`json:"name,omitempty"`
	Number	string		`json:"number,omitempty"`
}
var pokemon []Pokemon

func main(){

	pokemon = append(pokemon, Pokemon{ID: "1", Name: "Bulbasaur", Number: "1"})
	pokemon = append(pokemon, Pokemon{ID: "2", Name: "Ivysaur", Number: "2"})
	
	router := mux.NewRouter()
	
	router.HandleFunc("/pokemon", GetPokemon).Methods("GET")
	
	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetPokemon(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(pokemon)
}