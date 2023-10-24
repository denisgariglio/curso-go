package main

import "net/http"

func main() {
	// http.HandleFunc("/", BuscaCep)
	//exemplo igual ao anterior com função anonima
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world 2 !"))
	})
	http.ListenAndServe(":8080", nil)
}

func BuscaCep(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
