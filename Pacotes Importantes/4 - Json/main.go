package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"-"`
	Saldo int `json:"s"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	println(string(res))

	json.NewEncoder(os.Stdout).Encode(conta)

	jsonPuro := []byte(`{"n":2, "s":200}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		panic(err)
	}
	println(contaX.Saldo)
}