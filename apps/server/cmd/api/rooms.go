package main

import (
	"errors"
	"net/http"
)

func (app *application) createRoomHandler(w http.ResponseWriter, r *http.Request) {
	app.mu.Lock()
	defer app.mu.Unlock()
	var input struct {
		Name string `json:"name"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusBadRequest)
		return
	}

	newRoom := Room{
		Name:  input.Name,
		Users: [2]User{},
	}

	app.rooms[input.Name] = newRoom
	err = app.writeJSON(w, http.StatusOK, envelope{"rooms": app.rooms}, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}
}

func (app *application) listRoomsHandler(w http.ResponseWriter, r *http.Request) {
	err := app.writeJSON(w, http.StatusOK, envelope{"rooms": app.rooms}, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}
}

func (app *application) joinRoom(name string, user User) (string, error) {
	app.mu.Lock()
	defer app.mu.Unlock()

	room, ok := app.rooms[name]
	if !ok {
		return "", errors.New("Room does not exist")
	}

	if room.Users[0].Name == "" {
		room.Users[0] = user
	} else if room.Users[1].Name == "" {
		room.Users[1] = user
	} else {
		return "", errors.New("Room is full")
	}

	app.rooms[name] = room
	return "Successfully joined room:" + name, nil
}
