package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
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
	
	router.HandleFunc("/", DefaultRoute).Methods("GET")
	router.HandleFunc("/pokemon", GetPokemon).Methods("GET")
	router.HandleFunc("/pokemon/{id}", GetPokemonById).Methods("GET")
	router.HandleFunc("/pokemon/{id}", CreatePokemon).Methods("POST")
	router.HandleFunc("/pokemon/{id}", DeletePokemon).Methods("DELETE")
	
	log.Fatal(http.ListenAndServe(":8080", router))
}

func DefaultRoute(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode("Hello World")
}

func GetPokemon(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(pokemon)
}

func GetPokemonById(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//handle this
	}
	id -= 1;
	if(id < len(pokemon) && id >= 0){
		json.NewEncoder(w).Encode(pokemon[id])
	}else{
		json.NewEncoder(w).Encode(Pokemon{})
	}
}

func CreatePokemon(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
    var pkmn Pokemon
    _ = json.NewDecoder(r.Body).Decode(&pkmn)
    pkmn.ID = params["id"]
    pokemon = append(pokemon, pkmn)
    json.NewEncoder(w).Encode(pokemon)
}

func DeletePokemon(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
    for index, item := range pokemon {
        if item.ID == params["id"] {
            pokemon = append(pokemon[:index], pokemon[index+1:]...)
            break
        }
        json.NewEncoder(w).Encode(pokemon)
    }
}