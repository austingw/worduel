package main

import (
	"context"
	"errors"
	"net/http"
	"time"
	words "worduel/internal"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
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
		Name:        input.Name,
		Users:       [2]User{},
		CurrentWord: "",
	}

	app.rooms[input.Name] = newRoom
	err = app.writeJSON(w, http.StatusOK, envelope{"newRoom": newRoom}, nil)
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

func (app *application) joinRoom(name string, username string, conn *websocket.Conn) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	app.mu.Lock()
	defer app.mu.Unlock()

	room, ok := app.rooms[name]
	if !ok {
		return errors.New("Room does not exist")
	}

	if room.Users[0].Ws == nil {
		room.Users[0] = User{
			Name: username,
			Ws:   conn,
		}

		err := wsjson.Write(ctx, conn, envelope{"message": "Room " + name + "successfully created!"})
		if err != nil {
			app.logger.Error(err.Error())
		}
	} else if room.Users[1].Ws == nil {
		room.Users[1] = User{
			Name: username,
			Ws:   conn,
		}

		// get new word on player 2 join
		room.CurrentWord = words.NewWord()

		err := wsjson.Write(ctx, room.Users[0].Ws, envelope{"type": "start", "message": username + " joined room!"})
		if err != nil {
			app.logger.Error(err.Error())
		}
		err = wsjson.Write(ctx, conn, envelope{"type": "start", "message": "Room " + name + "successfully joined!"})
		if err != nil {
			app.logger.Error(err.Error())
		}
	} else {
		return errors.New("Room is full")
	}

	app.rooms[name] = room
	return nil
}

func (app *application) leaveRoom(name string, username string, conn *websocket.Conn) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	app.mu.Lock()
	defer app.mu.Unlock()

	// If user closes browser/loses connection, find info based on closed ws
	if name == "" {
		for roomName, room := range app.rooms {
			for _, user := range room.Users {
				if user.Ws == conn {
					name = roomName
					username = user.Name
				}
			}
		}
	}

	room, ok := app.rooms[name]
	if !ok {
		return errors.New("Room does not exist")
	}

	if conn == room.Users[1].Ws {
		room.Users[1] = User{}
		app.rooms[name] = room

		err := wsjson.Write(ctx, room.Users[0].Ws, envelope{"type": "end", "message": username + " left room."})
		if err != nil {
			app.logger.Error(err.Error())
		}

		return nil
	} else if conn == room.Users[0].Ws {
		if room.Users[1].Ws != nil {
			err := wsjson.Write(ctx, room.Users[1].Ws, envelope{"type": "hostEnd", "message": room.Users[0].Name + " left room. Closing room and return to lobby..."})
			if err != nil {
				app.logger.Error(err.Error())
			}
		}

		delete(app.rooms, name)
		return nil
	} else {
		return errors.New("User not in room???")
	}
}
