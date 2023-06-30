package main

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
	a := "X"
	println(a)
	x()
}

func x() {
	println(f, "Tipagem própria")
}
