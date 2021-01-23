package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handleHello)
	mux.HandleFunc("/goodbye", handleGoodbye)
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusMethodNotAllowed)
	})

	s := http.Server{
		Addr:    "localhost:80",
		Handler: mux,
	}

	fmt.Println(s.ListenAndServe())
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello HTTP"))
	if err != nil {
		fmt.Printf("ALARM %s", err)
	}
}

func handleGoodbye(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Goodbye HTTP"))
	if err != nil {
		fmt.Printf("ALARM %s", err)
	}
}
