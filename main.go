package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handleHello)
	mux.HandleFunc("/goodbye", handleGoodbye)
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusMethodNotAllowed)
	})
	mux.HandleFunc("/num", handleDouble)

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

func handleDouble(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		_, _ = w.Write([]byte("Method not Post"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.Write([]byte("Cant read request body " + err.Error()))
		return
	}

	number, err := strconv.Atoi(string(body))

	if err != nil {
		w.Write([]byte("Fail to parse numbers"))
		return
	}

	w.Write([]byte(strconv.Itoa(number * 2)))

}
