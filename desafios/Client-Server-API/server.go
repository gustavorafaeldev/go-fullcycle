package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Cotacao struct {
	USDBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func main() {
	http.HandleFunc("/cotacao", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if r.URL.Path != "/cotacao" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	ctx := r.Context()

	select {
	case <-time.After(time.Millisecond):
		// Processamento ok
		cotacao, error := getCotacao()
		if error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(cotacao)

		// salvar no banco
		ctxBanco, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
		defer cancel()
		saveCotacao(ctxBanco, db, cotacao)
	case <-ctx.Done():
		// Expiração da request
		log.Println("Falha na requisição")
		http.Error(w, "Falha na requisição", http.StatusRequestTimeout)
	}
}

func getCotacao() (*Cotacao, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		log.Println("Erro na requisição para Dolar 1")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Erro na requisição para Dolar 2")
	}

	defer res.Body.Close()

	body, error := ioutil.ReadAll(res.Body)
	if error != nil {
		return nil, nil
	}

	var response Cotacao
	error = json.Unmarshal(body, &response)
	if error != nil {
		return nil, nil
	}

	println(&response)
	return &response, nil
}

func saveCotacao(ctx context.Context, db *sql.DB, cotacao *Cotacao) error{
	select {
	case <-time.After(time.Millisecond):
		stmt, err := db.Prepare("INSERT INTO tabela_cotacao (code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, create_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			panic(err)
		}

		_, err = stmt.Exec(cotacao.USDBRL.Code, cotacao.USDBRL.Codein, cotacao.USDBRL.Name, cotacao.USDBRL.High, cotacao.USDBRL.Low, cotacao.USDBRL.VarBid, cotacao.USDBRL.PctChange, cotacao.USDBRL.Bid, cotacao.USDBRL.Ask, cotacao.USDBRL.Timestamp, cotacao.USDBRL.CreateDate)
		if err != nil {
			return err
		}
		return nil
	case <-ctx.Done():
		log.Println("Erro contexto salvamento")
		return nil
	}
}
