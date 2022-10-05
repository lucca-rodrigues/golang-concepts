package main

func main() {
	a := 1
	b := 2
	c := 0

	if a > b {
		println(a)
	} else {
		println(b)

	}

	if a > b || c < b{
		println("o valor não pode ser menor que", a)
	}

	switch a{
	case 1:
		println("O valor de a é 1")
	case 2:
		println("O valor de a é 2")
	default:
		println("Valor defalt aqui!")
	}
}