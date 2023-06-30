package main

// import "fmt"

func main() {
	var minhaVar interface{} = "Rafael Lira"

	println(minhaVar.(string))
	res, ok := minhaVar.(int)
	println(res, ok)
	res2 := minhaVar.(int)
	println(res2)
}
