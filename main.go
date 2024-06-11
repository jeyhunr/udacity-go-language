package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Customer struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}

var customersFile = "./customers.json"

func getCustomers() []Customer {
	file, err := os.ReadFile(customersFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return []Customer{}
	}
	var customers []Customer
	if err := json.Unmarshal(file, &customers); err != nil {
		fmt.Println("Error unmarshalling file:", err)
		return []Customer{}
	}
	return customers
}

func getCustomer(id string) Customer {
	customers := getCustomers()
	for _, customer := range customers {
		if customer.ID == id {
			return customer
		}
	}
	return Customer{}
}

func saveCustomers(customers []Customer) {
	data, err := json.Marshal(customers)
	if err != nil {
		fmt.Println("Error marshalling customers:", err)
		return
	}
	if err := os.WriteFile(customersFile, data, 0644); err != nil {
		fmt.Println("Error writing file:", err)
	}
}

func addCustomer(customer Customer) Customer {
	customers := getCustomers()
	customers = append(customers, customer)
	saveCustomers(customers)
	return customer
}

func updateCustomer(updatedCustomer Customer) Customer {
	customers := getCustomers()
	for i, customer := range customers {
		if customer.ID == updatedCustomer.ID {
			customers[i] = updatedCustomer
			saveCustomers(customers)
			return updatedCustomer
		}
	}
	return Customer{}
}

func deleteCustomer(id string) Customer {
	customers := getCustomers()
	var deletedCustomer Customer
	for i, customer := range customers {
		if customer.ID == id {
			deletedCustomer = customer
			customers = append(customers[:i], customers[i+1:]...)
			saveCustomers(customers)
			break
		}
	}
	return deletedCustomer
}

func main() {
	r := mux.NewRouter()

	// Serve the index.html file for the root route
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	}).Methods("GET")

	r.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		customers := getCustomers()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customers)
	}).Methods("GET")

	r.HandleFunc("/customer/{ID}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["ID"]
		customer := getCustomer(id)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customer)
	}).Methods("GET")

	r.HandleFunc("/customer", func(w http.ResponseWriter, r *http.Request) {
		var customer Customer
		if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		addedCustomer := addCustomer(customer)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(addedCustomer)
	}).Methods("POST")

	r.HandleFunc("/customer/{ID}", func(w http.ResponseWriter, r *http.Request) {
		var customer Customer
		if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		updatedCustomer := updateCustomer(customer)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedCustomer)
	}).Methods("PUT")

	r.HandleFunc("/customer/{ID}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["ID"]
		deletedCustomer := deleteCustomer(id)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(deletedCustomer)
	}).Methods("DELETE")

	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", r)
}
