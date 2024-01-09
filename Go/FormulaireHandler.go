package HangMan_Web

import (
	"html/template"
	"log"
	"net/http"

	hangman "github.com/Ssnakyx/HangMan____"
)

func FormulaireHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Erreur", http.StatusInternalServerError)
			return
		}

		lettre := r.FormValue("Lettre")
		foundLetters = append(foundLetters, lettre)

		data := HangManData{
			WordToGuess:       currentWord,
			AttemptsRemaining: 6,
			GameStage:         []string{},
			Lettre:            lettre,
		}

		data.DisplayedWord = hangman.DisplayWord(currentWord, foundLetters)

		if hangman.IsWordGuessed(currentWord, foundLetters) {

			http.Redirect(w, r, "/victoire", http.StatusFound)
			return
		}

		template, err := template.ParseFiles("HTMLL/forms.html", "HTMLL/footer.html", "HTMLL/header.html")
		if err != nil {
			log.Fatal(err)
		}

		template.Execute(w, data)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
