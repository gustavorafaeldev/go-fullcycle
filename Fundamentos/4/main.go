package main

import "fmt"

const a = "Hello, World"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Iniciando com valor"
	e float64 = 10.123456
	f ID      = 1
)

func main() {
	fmt.Printf("O tipo d E é %T", f)
}
