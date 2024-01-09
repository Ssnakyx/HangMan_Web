package HangMan_Web

import (
	"html/template"
	"log"
	"net/http"
)

func DefeatePage(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("HTMLL/defaite.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}
