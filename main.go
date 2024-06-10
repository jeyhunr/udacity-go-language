package main

import (
	"encoding/json"
	"fmt"
	"os"
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

func main() {
	customers := getCustomers()
	customer := getCustomer("2")

	newCustomer := Customer{ID: "12345", Name: "Jeyhun Rahimli", Role: "Software Engineer", Email: "mail@rahimli.net", Phone: "0123456", Contacted: true}
	fmt.Println(customers)
	fmt.Println(customer)

	addCustomer(newCustomer)
}
