package HangMan_Web

import (
	"log"
	"net/http"
)

var (
	currentWord  string
	foundLetters []string
	WordToGuess  string
)

type HangManData struct {
	WordToGuess       string
	DisplayedWord     string
	AttemptsRemaining int
	GameStage         []string
	Lettre            string
	Score             int
}

func Serveur() {
	http.HandleFunc("/levels", LevelsPage)
	http.HandleFunc("/", Home)
	http.HandleFunc("/startgame", StartGamePage)
	http.HandleFunc("/hangman", FormulaireHandler)
	http.HandleFunc("/victoire", VictoirePage)
	http.HandleFunc("/defaite", DefeatePage)
	fs := http.FileServer(http.Dir("CSS/"))
	http.Handle("/CSS/", http.StripPrefix("/CSS", fs))
	log.Println("Serveur allumé")
	http.ListenAndServe(":8080", nil)
}
