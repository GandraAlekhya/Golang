package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var profiles []profile = []Profile{}

type User struct {
	firstName string `json:"firstName"`
	lastName  string `json:"lastName"`
}

type Profile struct {
	AccNum    int    `json:accNum`
	AccType   string `json:"accType"`
	AccHolder User   `json:"accHolder"`
}

func addMember(q http.ResponseWriter, r *http.Request) {
	var newProfile Profile
	json.NewDecoder(r.Body).Decode(&newProfile)

	w.Header().Set("Content-Type", application/json)

	profiles = append(profiles, newProfile)
	json.NewEncoder(q).Encode(profiles)

}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/profiles", addMember).Methods("POST")

	http.ListenAndServe(":5000", router)

}
