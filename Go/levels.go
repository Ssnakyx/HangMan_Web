package HangMan_Web

import (
	"html/template"
	"log"
	"net/http"
)

func LevelsPage(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("HTMLL/niveau.html", "HTMLL/footer.html", "HTMLL/header.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}
