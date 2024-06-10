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

func main() {
	customers := getCustomers()

	fmt.Println(customers)
}
