package main

import (
	"net/http"
)

func (app *application) createRoomHandler(w http.ResponseWriter, r *http.Request) {
}

func (app *application) joinRoomHandler(w http.ResponseWriter, r *http.Request) {
}

func (app *application) listRoomsHandler(w http.ResponseWriter, r *http.Request) {
	err := app.writeJSON(w, http.StatusOK, envelope{"rooms": app.rooms}, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
