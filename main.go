package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"pokemon-api/database"
)

func getAllPokemons(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(database.PokemonDb)
}

func addNewPokemon(w http.ResponseWriter, r *http.Request){
	reqBody, _ := ioutil.ReadAll(r.Body)

	var newPokemon database.Pokemon
	json.Unmarshal(reqBody, &newPokemon)

	database.PokemonDb = append(database.PokemonDb, newPokemon)

	json.NewEncoder(w).Encode(newPokemon)
}

func handleRequests() {

	port := os.Getenv("PORT")

	if port==""{
		port = "80"
	}

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Use(commonMiddleware)
	myRouter.HandleFunc("/pokemons/get", getAllPokemons).Methods("GET")
	myRouter.HandleFunc("/pokemons/add", addNewPokemon).Methods("POST")

	log.Fatal(http.ListenAndServe(":"+port, myRouter))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	fmt.Println("Pokemon Rest API")
	handleRequests()
}
