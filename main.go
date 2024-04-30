package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var profiles []Profile = []Profile{}

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}
type Profile struct {
	Department  string `json:"department"`
	Designation string `json:"designation"`
	Employee    User   `json:"employee"`
}

func addItem(h http.ResponseWriter, r *http.Request) {
	var newProfile Profile
	json.NewDecoder(r.Body).Decode(&newProfile)

	h.Header().Set("Content-Type", "application/json")

	profiles = append(profiles, newProfile)

	json.NewEncoder(h).Encode(profiles)
}

func getAllProfiles(h http.ResponseWriter, r *http.Request) {
	h.Header().Set("Content-Type", "application/json")

	json.NewEncoder(h).Encode(profiles)
}

func getProfiles(h http.ResponseWriter, r *http.Request) {

	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		h.WriteHeader(400)
		h.Write([]byte("id could not be converted to int"))
		return
	}

	if id >= len(profiles) {
		h.WriteHeader(404)
		h.Write([]byte("no profiles found with specified ID"))
		return
	}
	profile := profiles[id]

	h.Header().Set("Content-Type", "application/json")
	json.NewEncoder(h).Encode(profile)
}

func updateProfiles(h http.ResponseWriter, r *http.Request) {

	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		h.WriteHeader(400)
		h.Write([]byte("id could not be converted to int"))
		return
	}

	if id >= len(profiles) {
		h.WriteHeader(404)
		h.Write([]byte("no profiles found with specified ID"))
		return
	}

	var updateProfile Profile

	json.NewDecoder(r.Body).Decode(&updateProfile)

	profiles[id] = updateProfile
	h.Header().Set("Content-Type", "application/json")
	json.NewEncoder(h).Encode(updateProfile)
}

func deleteProfiles(h http.ResponseWriter, r *http.Request) {

	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		h.WriteHeader(400)
		h.Write([]byte("id could not be converted to int"))
		return
	}

	if id >= len(profiles) {
		h.WriteHeader(404)
		h.Write([]byte("no profiles found with specified ID"))
		return
	}

	profiles = append(profiles[:id], profiles[id+1:]...)

	h.WriteHeader(200)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/profiles", addItem).Methods("POST")

	router.HandleFunc("/profiles", getAllProfiles).Methods("GET")

	router.HandleFunc("/profiles/{id}", getProfiles).Methods("GET")

	router.HandleFunc("/profiles/{id}", updateProfiles).Methods("PUT")

	router.HandleFunc("/profiles/{id}", deleteProfiles).Methods("DELETE")

	http.ListenAndServe(":5000", router)
}
