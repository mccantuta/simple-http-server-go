package main

import (
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Id   string
	Name string
}

func customerHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "Testing GET")
	case http.MethodPost:
		fmt.Fprintf(w, "Testing POST")
	default:
		fmt.Fprintf(w, "Testing")
	}

}

func main() {
	http.HandleFunc("/customer", customerHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
