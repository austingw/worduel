package main

import (
	"math/rand"
	"net/http"
	words "worduel/internal"
)

func (app *application) checkAnswer(w http.ResponseWriter, r *http.Request) {
}

func (app *application) newWord() string {
	return words.List[rand.Intn(len(words.List))]
}
