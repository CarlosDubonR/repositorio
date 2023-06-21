package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola Mundoooooo")
	})
	http.HandleFunc("/sig", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "comoooo estaaaan jeje")
	})
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
