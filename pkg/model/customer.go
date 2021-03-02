package model

import (
	"encoding/json"
	"log"
)

type Customer struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func GetCustomer() ([]byte, error) {
	customerList := []Customer{
		Customer{"001", "Customer 1"},
		Customer{"002", "Customer 2"},
		Customer{"003", "Customer 3"},
	}
	return json.MarshalIndent(customerList, "", "")
}

func PostCustomer(data []byte) string {
	var customer Customer
	e := json.Unmarshal(data, &customer)
	if e != nil {
		log.Fatal(e)
	}
	log.Println(customer)
	return "Customer saved"
}
