package main

import "fmt"

type Conta struct {
	saldo float64
}

func NewConta() *Conta{
	return &Conta{saldo: 0}
}

func (c *Conta) simular(valor float64) float64{
	c.saldo += valor
	fmt.Println(c.saldo)
	return c.saldo
}

func main() {
	conta := Conta{saldo: 100}
	conta.simular(200)
	fmt.Println(conta.saldo)
}