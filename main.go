package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Auto struct {
	ID      uint   `json:"id"`
	Brand   string `json:"brand"`
	Model   string `json:"model"`
	Price   uint   `json:"price"`
	Status  uint   `json:"status"`
	Mileage uint   `json:"mileage"`
}

var StatusMap = map[string]uint{
	"In transit":          0,
	"In stock":            1,
	"Sold":                2,
	"Withdrawn from sale": 3,
}

var autos []Auto

func main() {
	r := mux.NewRouter()

	autos = append(autos, Auto{ID: 1, Brand: "BMW", Model: "M5", Price: 2440000, Status: StatusMap["In transit"], Mileage: 5000})
	autos = append(autos, Auto{ID: 2, Brand: "Mazda", Model: "RX-7", Price: 1540000, Status: StatusMap["Sold"], Mileage: 15000})
	autos = append(autos, Auto{ID: 3, Brand: "Lada", Model: "Vesta Cross", Price: 861000, Status: StatusMap["In stock"], Mileage: 0})

	r.HandleFunc("/autos", getAutos).Methods("GET")
	r.HandleFunc("/autos/{id}", getAuto).Methods("GET")
	r.HandleFunc("/autos", createAuto).Methods("POST")
	r.HandleFunc("/autos/{id}", updateAuto).Methods("PUT")
	r.HandleFunc("/autos/{id}", deleteAuto).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func getAutos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(autos)
}

func getAuto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 8)
	for _, item := range autos {
		if item.ID == uint(id) {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Auto{})
}

func createAuto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var auto Auto
	_ = json.NewDecoder(r.Body).Decode(&auto)
	auto.ID = uint(rand.Intn(1000000))
	autos = append(autos, auto)
	json.NewEncoder(w).Encode(auto)
}

func updateAuto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 8)
	for index, item := range autos {
		if item.ID == uint(id) {
			autos = append(autos[:index], autos[index+1:]...)
			var auto Auto
			_ = json.NewDecoder(r.Body).Decode(&auto)
			auto.ID = uint(id)
			autos = append(autos, auto)
			json.NewEncoder(w).Encode(auto)
			return
		}
	}
	json.NewEncoder(w).Encode(autos)
}

func deleteAuto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 8)
	for index, item := range autos {
		if item.ID == uint(id) {
			autos = append(autos[:index], autos[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(autos)
}
