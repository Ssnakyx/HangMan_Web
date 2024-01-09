package HangMan_Web

import (
	"html/template"
	"log"
	"net/http"
)

func VictoirePage(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("HTMLL/victoire.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}
