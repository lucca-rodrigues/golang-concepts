package main

import "fmt"

func main(){
	s :=[]int{2,4,6,8,10}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s) // get all icons

	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0]) // no get itens

	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4]) // get itens finishend to position four

	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:]) // get itens started to third position
}