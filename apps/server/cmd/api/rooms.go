package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) createRoomHandler(w http.ResponseWriter, r *http.Request) {
}

func (app *application) joinRoomHandler(w http.ResponseWriter, r *http.Request) {
}

func (app *application) listRoomsHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"hello": "world",
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
