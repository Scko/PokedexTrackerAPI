package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	//"io/ioutil"
	//"fmt"
)

type PokeAPI struct {
	Count int
	Next int
	Previous int
	Results []Result
}

type Result struct {
	Name string
	Url string
}

type Poke struct {
	Id int
	Name string
	Height int
	Weight int
	Sprites []Sprite
}

type Sprite struct{
	Back_Default string
	Back_Female string
	Back_Shiny string
	Back_Shiny_Female string
	Front_Default string
	Front_Female string
	Front_Shiny string
	Front_Shiny_Female string
}

type Pokemon struct{
	ID		string		`json:"id,omitempty"`
	Name	string		`json:"name,omitempty"`
	Number	string		`json:"number,omitempty"`
}
var pokemon []Pokemon
var pokes []Poke

func main(){

	pokemon = append(pokemon, Pokemon{ID: "1", Name: "Bulbasaur", Number: "1"})
	pokemon = append(pokemon, Pokemon{ID: "2", Name: "Ivysaur", Number: "2"})
	
	router := mux.NewRouter()
	
	router.HandleFunc("/", DefaultRoute).Methods("GET")
	router.HandleFunc("/pokemon", GetPokemon).Methods("GET")
	router.HandleFunc("/pokemon/{id}", GetPokemonById).Methods("GET")
	router.HandleFunc("/pokemon/{id}", CreatePokemon).Methods("POST")
	router.HandleFunc("/pokemon/{id}", DeletePokemon).Methods("DELETE")
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("listening on port %v\n", port)

	log.Fatal(http.ListenAndServe(":"+port, router))
	
	
}

func DefaultRoute(w http.ResponseWriter, r *http.Request){
	//Get all the pokemon urls
	response, err := http.Get("https://pokeapi.co/api/v2/pokemon/")
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	var data PokeAPI
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		log.Println(err)
	}
	//json.NewEncoder(w).Encode(data)
	
	//For each result make an individual request to get that pokemons data
	for _, item := range data.Results{
		response, err := http.Get(item.Url)
		if err != nil {
			json.NewEncoder(w).Encode(err)
		}
		var p Poke
		if err := json.NewDecoder(response.Body).Decode(&p); err != nil {
			log.Println(err)
		}
		pokes = append(pokes, p)
	}
	
	//Return the data of all
	json.NewEncoder(w).Encode(pokes)
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