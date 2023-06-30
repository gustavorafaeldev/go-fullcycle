package main

func soma(a, b *int) int {
	*a = 50
	*b = 50
	return *a + *b
}

// Variaveis mutaveis geralmente se usam ponteiro, para alterar o valor da variavel naquele endereço.

func main() {
	minhaVar1 := 10
	minhaVar2 := 20
	soma(&minhaVar1, &minhaVar2)

	println(minhaVar1, "Esse é o valor")
	println(&minhaVar1, "Esse é o endereço")

	println(minhaVar1)
	println(minhaVar2)
}