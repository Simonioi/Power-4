package main

import (
	"html/template"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	data := map[string]string{
		"Titre":   "Bienvenue sur ma page Go + HTML",
		"Message": "Ceci est un exemple de liaison entre Go et HTML.",
	}
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
