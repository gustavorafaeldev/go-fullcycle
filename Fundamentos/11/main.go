package main

import "fmt"

// Qualquer que chamar o método desativar automaticamente está implementando essa interface
// Interface do Go só pode conter métodos, não aceita atributos
type Pessoa interface {
	Desativar()
}

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome    string
	Idade   int
	Ativo   bool
	Address Endereco
}

type Empresa struct {
	Nome string
}

// Esse tipo de função é como se fosse um método do struct
func (e Empresa) Desativar() {
	fmt.Println("Empresa desativada")
}

// Esse tipo de função é como se fosse um método do struct
func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	Rafael := Cliente{
		Nome:  "Rafael",
		Idade: 30,
		Ativo: true,
	}
	Rafael.Ativo = false
	Rafael.Address.Cidade = "Sarandi"

	minhaEmpresa := Empresa{}

	Desativacao(minhaEmpresa)
	Desativacao(Rafael)


	//fmt.Printf("Nome: %s, Idade %d, Ativo: %t", Rafael.Nome, Rafael.Idade, Rafael.Ativo)
	//fmt.Println(Rafael)
}