package main

import "fmt"

func main()  {
	sallary := map[string]int{"Lucas": 1000, "João": 1000, "Novell": 1000}

	//fmt.Println(sallary["Lucas"])
	delete(sallary, "João") // Remove item
	sallary["Pedro"] = 5000 // add new icon
	fmt.Println(sallary) // map[Lucas:1000 Novell:1000 Pedro:5000]
}