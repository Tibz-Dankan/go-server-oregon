package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Person struct {
	Name          string `json:"name"`
	Continent     string `json:"continent"`
	FavoriteSport string `json:"favorite_sport"`
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("GO server Oregon started on port 8080")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	people := generatePeople(40)
	jsonData, err := json.Marshal(people)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Getting random people...")
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func generatePeople(count int) []Person {
	rand.Seed(time.Now().UnixNano())
	names := []string{"John", "Alice", "Bob", "Emma", "Michael", "Sophia", "William", "Olivia", "James", "Emily"}
	continents := []string{"Asia", "Europe", "Africa", "North America", "South America", "Australia", "Antarctica"}
	sports := []string{"Football", "Basketball", "Tennis", "Cricket", "Golf", "Baseball", "Rugby", "Hockey", "Soccer"}

	people := make([]Person, count)
	for i := 0; i < count; i++ {
		people[i] = Person{
			Name:          names[rand.Intn(len(names))],
			Continent:     continents[rand.Intn(len(continents))],
			FavoriteSport: sports[rand.Intn(len(sports))],
		}
	}
	return people
}
