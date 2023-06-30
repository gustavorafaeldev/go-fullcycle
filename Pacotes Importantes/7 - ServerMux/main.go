package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	// mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
	// 	w.Write([]byte("Hello, World"))
	// })
	mux.HandleFunc("/", HomeHandler)
	// mux.Handle("/blog", blog{title: "My Blog"})

	http.ListenAndServe(":8080", mux)

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Wesley"))
	})
	http.ListenAndServe(":8081", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

type blog struct {
	title string
}

func (b blog) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}