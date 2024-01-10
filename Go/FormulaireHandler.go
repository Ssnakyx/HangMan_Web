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
		if len(foundLetters) > 10 {
			currentWord = pickNewWord()
			foundLetters = []string{}
			http.Redirect(w, r, "/defaite?currentWord="+currentWord, http.StatusFound)
			return
		}

		data := HangManData{
			WordToGuess:       currentWord,
			AttemptsRemaining: 10 - len(foundLetters),
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

func pickNewWord() string {
	words, err := hangman.ReadWordsFromFile("words/words.txt")
	if err != nil {
		log.Fatal(err)
	}

	return hangman.GetRandomWordFromList(words)
}
func StartGamePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Erreur", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/hangman", http.StatusFound)
		return
	}
	template, err := template.ParseFiles("HTMLL/startgame.html", "HTMLL/footer.html", "HTMLL/header.html")
	if err != nil {
		http.Error(w, "Nul", http.StatusInternalServerError)
		return
	}
	template.Execute(w, nil)
}
