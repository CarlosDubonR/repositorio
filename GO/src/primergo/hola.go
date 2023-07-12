package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

/*
	func Hola(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>texto servidor pag principal jsjsjsjsj  </h1>")
	}
*/
func sig(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>siguiente direccion  </h1>")
}

type mensaje struct {
	msg string
}

func (m mensaje) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.msg)
}

func main() {
	msg := mensaje{msg: "texto desde un handle:  --"}
	estatic := http.FileServer(http.Dir("estatico"))
	mux := http.NewServeMux()
	mux.Handle("/", estatic)
	mux.HandleFunc("/sig", sig)
	mux.Handle("/handle", msg)
	server := &http.Server{
		Addr:           ":8080",          //especificamos puerto serie
		Handler:        mux,              //especifica el server mux, ya sea por defecto o el nuestro
		ReadTimeout:    10 * time.Second, //timepos en los que entran pedidod y salen respuestas
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, //numero de bytes
	}
	log.Println("corriendo servidor propio")
	log.Fatal(server.ListenAndServe()) //en listenAndServe no se mandan parametros por que el puerto y el server mux ya los defini

	//http.ListenAndServe(":8080", mux) //el nill es el server mux por defecto y el mux es algo que yo cree
}
