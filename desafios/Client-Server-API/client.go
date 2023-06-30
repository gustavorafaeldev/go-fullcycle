package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"
)


type CotacaoResponse struct {
	USDBRL struct {
		Bid        string `json:"bid"`

	} `json:"USDBRL"`
}


func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	response, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var data CotacaoResponse
	err = json.Unmarshal(response, &data)
	if err != nil {
		panic(err)
	}
	saveFile(data.USDBRL.Bid)
}

func saveFile(valor string) {
	f, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}

	f.Write([]byte("Dólar: "+ valor))
}
