package main

import "fmt"

type MyNumber int

type Number interface{
	~int | float64
}

func Soma[T Number] (m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a, b T) bool {
	if a == b {
		return true
	}
	return false
}


func main() {
	m := map[string]int{"Rafael": 10000, "João": 2000, "Maria": 3000}
	m2 := map[string]float64{"Rafael": 10000.20, "João": 2000.55, "Maria": 3000.30}
	m3 := map[string]MyNumber{"Rafael": 10000, "João": 2000, "Maria": 3000}

	println(Soma(m))
	fmt.Print(Soma(m2))
	fmt.Print(Soma(m3))
	println(Compara(10, 10))
}
