package main

import "fmt"

type Peaple interface  {
	Create()
}

func Create(){
	fmt.Printf("User created successfully")
}
func main() {}