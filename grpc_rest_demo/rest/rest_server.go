package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	user := User{ID: id, Name: fmt.Sprintf("User %s", id)}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func main() {
	http.HandleFunc("/user", getUserHandler)
	log.Println("REST server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
