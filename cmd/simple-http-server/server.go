package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mccantuta/simple-http-server/pkg/model"
)

func customerHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		data, e := model.GetCustomer()
		if e != nil {
			log.Fatal("Error getting Customers")
		}
		w.Write(data)
	case http.MethodPost:
		data, e := ioutil.ReadAll(r.Body)
		if e != nil {
			log.Fatal("Error reading the request body")
		}
		result := model.PostCustomer(data)
		w.Write([]byte(result))
	default:
		fmt.Fprintf(w, "Not supported operation")
	}
}

func main() {
	http.HandleFunc("/customer", customerHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
