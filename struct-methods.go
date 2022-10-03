package main

import "fmt"


type Customer struct {
	Name string
	Active bool
}


func (c Customer) Disable() {
	c.Active = false

	fmt.Printf("O cliente foi desativado")
}

func main(){
	customer := Customer{
		Name:   "Lucas",
		Active: true,
	}

	customer.Disable()
}