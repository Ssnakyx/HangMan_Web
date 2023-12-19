package main

import (
	"log"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("HTMLL/forms.html", "HTMLL/footer.html", "HTMLL/header.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}

func FormulaireHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Erreur", http.StatusInternalServerError)
			return
		}

		lettre := r.FormValue("Lettre")
		log.Print(lettre)

		data := struct {
			Lettre string
		}{Lettre: lettre}

		template, err := template.ParseFiles("HTMLL/forms.html", "HTMLL/footer.html", "HTMLL/header.html")
		if err != nil {
			log.Fatal(err)
		}

		template.Execute(w, data)
		return
	}

	http.Redirect(w, r, "http://localhost:8080", http.StatusFound)
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/hangman", FormulaireHandler)

	fs := http.FileServer(http.Dir("CSS/"))
	http.Handle("/CSS/", http.StripPrefix("/CSS", fs))
	log.Println("Serveur allumé")
	http.ListenAndServe(":8080", nil)
}
