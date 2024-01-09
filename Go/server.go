package HangMan_Web

import (
	"log"
	"net/http"
)

var (
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

func Serveur() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/hangman", FormulaireHandler)
	http.HandleFunc("/victoire", VictoirePage)
	http.HandleFunc("/defeate", DefeatePage)
	fs := http.FileServer(http.Dir("CSS/"))
	http.Handle("/CSS/", http.StripPrefix("/CSS", fs))
	log.Println("Serveur allumé")
	http.ListenAndServe(":8080", nil)
}
