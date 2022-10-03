package main

import "fmt"

type Address struct {
	Address string
	Number int
}

type Customer struct {
	Name string
	Age int
	Address
}

func main(){
	newCustomer := Customer{Name: "Lucas", Age: 27}

	fmt.Printf("Nome: %s, Idade: %d", newCustomer.Name, newCustomer.Age)
}