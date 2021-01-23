package main

import (
	"fmt"
	"net/http"
)

func main() {
	s := http.Server{
		Addr:    "localhost:80",
		Handler: http.HandlerFunc(handleHello),
	}

	fmt.Println(s.ListenAndServe())
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello HTTP"))
	if err != nil {
		fmt.Printf("ALARM %s", err)
	}
}
