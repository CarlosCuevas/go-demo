package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	//params := mux.Vars(req)
	m1 := Matchup{{Name: "Giants", Score: 123}, {Name: "Patriots", Score: 100}}
	m2 := Matchup{{Name: "Jets", Score: 53}, {Name: "Cardinals", Score: 90}}
	m3 := Matchup{{Name: "Jags", Score: 83}, {Name: "Broncos", Score: 38}}
	m4 := Matchup{{Name: "Bills", Score: 73}, {Name: "Browns", Score: 104}}
	m5 := Matchup{{Name: "Chargers", Score: 84}, {Name: "Falcons", Score: 134}}
	matchups := []Matchup{m1, m2, m3, m4, m5}
	teams := QuantifyLuck(matchups)
	json.NewEncoder(w).Encode(teams)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/luck", GetPeopleEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
