package main

import "fmt"

func main()  {
	test := sum(1, 2)

	fmt.Println(test)
}

func sum(a int, b int)int{
	return a + b
}
func sum2(a, b int)int{
	return a + b
}

func sum3(a, b int)(int, bool){
	if a+b >= 50 {
		return a +b, true
	}
	return a + b, false
}