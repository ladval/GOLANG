package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var plantillas = template.Must(template.ParseGlob("templates/*"))

func main() {
	http.HandleFunc("/", Index)
	fmt.Println("Servidor corriendo...")
	http.ListenAndServe(":8080", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Hola mundo")
	plantillas.ExecuteTemplate(w, "index", nil)
}
