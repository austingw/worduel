package main

import (
	"net/http"
	words "worduel/internal"
)

func (app *application) checkAnswer(w http.ResponseWriter, r *http.Request) {
	words.NewWord()
}
