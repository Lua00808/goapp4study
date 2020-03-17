// TOOD: Make makefile

package main

import (
	"log"
	"net/http"
	"text/template"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/index.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
