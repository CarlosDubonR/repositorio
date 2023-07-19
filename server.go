package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

type Prueba struct {
	Title  string
	Numero int
}

type Pruebas []Prueba

/*const datos = {
	nombre: 'carlos',
	apellido: 'dubon'
}*/

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //utiliza un server mux interno
		prueba := Prueba{"dato en json", 63}
		json.NewEncoder(w).Encode(prueba)
		pruebas := Pruebas{
			Prueba{"dato en json", 456},
			Prueba{"para poder ", 343},
			Prueba{"conectar frontend", 673},
		}
		json.NewEncoder(w).Encode(pruebas)
	})

	mux.HandleFunc("/sig", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>codigo pruebas</h1>")
	})
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8080", handler)

	/*	msg := mensaje{msg: "texto desde un handle:  --"}
		//estatic := http.FileServer(http.Dir("estatico"))
		mux := http.NewServeMux()

		var nombres [5]string

		nombres[0] = "Ana"
		nombres[1] = "José"
		nombres[2] = "Daniel"
		nombres[3] = "María"
		nombres[4] = "Carlos"

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			prueba := Prueba{"prueba de datos json", 30}
			json.NewEncoder(w).Encode(prueba)
		})
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

		http.ListenAndServe(":8080", nil) //el nill es el server mux por defecto y el mux es algo que yo cree
	*/
}
