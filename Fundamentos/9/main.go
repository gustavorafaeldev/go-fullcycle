package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(1, 3, 4, 5, 6, 7, 5666, 54, 78, 56, 87, 567, 5, 654, 645))
}

func sum(numeros ...int) int {
	total := 0
	fmt.Println(numeros)
	for _, numero := range numeros {
		total += numero
	}
	return total
}
