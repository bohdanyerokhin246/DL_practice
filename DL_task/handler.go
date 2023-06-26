package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// TASK
func createHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	address := r.URL.Query().Get("address")

	err := insertRecord(name, address)
	if err != nil {
		log.Println("Error creating :", err)
		return
	}

	fmt.Fprintln(w, "Data is successfully added")
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	addresses, err := getRecordsByName(name)
	if err != nil {
		log.Println("Error SELECT from DB:", err)
		return
	}

	jsonData, err := json.Marshal(addresses)
	if err != nil {
		log.Println("Error writing JSON:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	address := r.URL.Query().Get("address")

	err := updateRecord(id, name, address)
	if err != nil {
		log.Println("Error UPDATE in DB:", err)
		return
	}

	fmt.Fprintln(w, "Data is successfully updated")
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	err := deleteRecord(id)
	if err != nil {
		log.Println("Error DELETE from DB:", err)
		return
	}

	fmt.Fprintln(w, "Data is successfully deleted")
}
