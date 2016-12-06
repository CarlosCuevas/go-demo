package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

func PostLuckEndpoint(w http.ResponseWriter, req *http.Request) {

	// m1 := Matchup{{Name: "Giants", Score: 123}, {Name: "Patriots", Score: 100}}
	// m2 := Matchup{{Name: "Jets", Score: 53}, {Name: "Cardinals", Score: 90}}
	// m3 := Matchup{{Name: "Jags", Score: 83}, {Name: "Broncos", Score: 38}}
	// m4 := Matchup{{Name: "Bills", Score: 73}, {Name: "Browns", Score: 104}}
	// m5 := Matchup{{Name: "Chargers", Score: 84}, {Name: "Falcons", Score: 134}}
	// matchups := []Matchup{m1, m2, m3, m4, m5}

	dec := json.NewDecoder(req.Body)
	for {
		var m Matchups
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(QuantifyLuck(m))
	}

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/luck", PostLuckEndpoint).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
