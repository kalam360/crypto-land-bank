package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world!")
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Kalam", "Dhaka", "1215"},
		{"Piash", "Dhaka", "1206"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func getOneCustomer(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "application/json")

	customer := Customer{"Saad", "Dhaka", "1214"}

	json.NewEncoder(w).Encode(customer)
	
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/get-customers", getCustomers)
	http.HandleFunc("/get-one-customer", getOneCustomer)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
