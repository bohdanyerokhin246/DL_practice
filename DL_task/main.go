package main

import "net/http"

// TASK
func main() {
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/update", updateHandler)
	http.HandleFunc("/delete", deleteHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
