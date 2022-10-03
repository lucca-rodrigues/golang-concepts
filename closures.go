package main

import "fmt"

func main(){
	total := func() int {
		return 1
	}()
	fmt.Println(total)
}