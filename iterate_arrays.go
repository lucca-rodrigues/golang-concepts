package main

import "fmt"

func main(){
	var myArray[3]int
	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30

	for index, item := range myArray {
		fmt.Printf("O valor do index é %d e o valor é  %d\n", index, item)
	}
	fmt.Println(myArray[0]) // get first item
}