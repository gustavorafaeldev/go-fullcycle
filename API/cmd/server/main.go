package main

import "github.com/gustavorafaeldev/curso-go/APIS/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}