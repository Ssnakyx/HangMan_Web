package HangMan_Web

import (
	"html/template"
	"log"
	"net/http"

	hangman "github.com/Ssnakyx/HangMan____"
)

func Home(w http.ResponseWriter, r *http.Request) {

	if currentWord == "" {
		words, err := hangman.ReadWordsFromFile("words/words.txt")
		if err != nil {
			log.Fatal(err)
		}

		currentWord = hangman.GetRandomWordFromList(words)
	}

	data := HangManData{
		WordToGuess:       currentWord,
		AttemptsRemaining: 10,
		GameStage:         []string{},
		Lettre:            "",
		DisplayedWord:     hangman.DisplayWord(currentWord, foundLetters),
	}

	template, err := template.ParseFiles("HTMLL/forms.html", "HTMLL/footer.html", "HTMLL/header.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, data)
}
