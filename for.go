package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	numbers := []string{"1", "2", "3"}

	for k, v := range numbers {
		println(k, v)
	}

	for i := 0; i < 10; i++ {
		println(i)
		i++
	}

}