package main

import (
	"fmt"
	"github.com/gustavorafaeldev/curso-go/7-Packaging/1/math"
)

func main() { 
	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add()) 
}