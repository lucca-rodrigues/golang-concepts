package main

import "fmt"

func main(){
	var myArray[3]int
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 10

	fmt.Println("Array", myArray)
	fmt.Println(len(myArray) -1)
	fmt.Println(myArray[len(myArray) -1]) // get last item
	fmt.Println(myArray[0]) // get first item
}