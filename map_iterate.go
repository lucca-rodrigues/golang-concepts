package main

import "fmt"



func main()  {
	sallary := map[string]int{"Lucas": 1000, "João": 1000, "Novell": 1000}

	for name, value := range sallary {
		fmt.Printf("O Salário de %s é %d\n", name, value)
	}
}