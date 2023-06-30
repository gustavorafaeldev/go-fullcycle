package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gustavorafaeldev/curso-go/matematica"
)

func main() {
	s := matematica.Soma(10, 20)
	carro := matematica.Carro{Marca: "Fiat"}
	fmt.Printf("Resultado: %v\n", s)
	fmt.Println(matematica.A)
	fmt.Println(carro)

	fmt.Println(uuid.New())
}

