package main

import (
	"html/template"
	"log"
	"net/http"
	"sync"

	hangman "github.com/Ssnakyx/HangMan____"
)

var (
	mu           sync.Mutex
	currentWord  string
	foundLetters []string
)

type HangManData struct {
	WordToGuess       string
	DisplayedWord     string
	AttemptsRemaining int
	GameStage         []string
	Lettre            string
}

func Home(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	if currentWord == "" {
		words, err := hangman.ReadWordsFromFile("words.txt")
		if err != nil {
			log.Fatal(err)
		}

		currentWord = hangman.GetRandomWordFromList(words)
	}

	data := HangManData{
		WordToGuess:       currentWord,
		AttemptsRemaining: 6,
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

func FormulaireHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

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

func VictoirePage(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("HTMLL/victoire.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/hangman", FormulaireHandler)
	http.HandleFunc("/victoire", VictoirePage)
	fs := http.FileServer(http.Dir("CSS/"))
	http.Handle("/CSS/", http.StripPrefix("/CSS", fs))
	log.Println("Serveur allumé")
	http.ListenAndServe(":8080", nil)
}
