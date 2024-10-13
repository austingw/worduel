package main

import (
	"context"
	"errors"
	"net/http"
	"time"

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
		Name:  input.Name,
		Users: [2]User{},
		// TODO: add logic to randomly pull word from list
		CurrentWord: "tares",
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

func (app *application) joinRoom(name string, username string, conn *websocket.Conn) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	app.mu.Lock()
	defer app.mu.Unlock()

	room, ok := app.rooms[name]
	if !ok {
		return "", errors.New("Room does not exist")
	}

	if room.Users[0].Name == "" {
		room.Users[0] = User{
			Name: username,
			Ws:   conn,
		}

		err := wsjson.Write(ctx, conn, envelope{"message": "Room " + name + "successfully created!"})
		if err != nil {
			app.logger.Error(err.Error())
		}
	} else if room.Users[1].Name == "" {
		room.Users[1] = User{
			Name: username,
			Ws:   conn,
		}

		err := wsjson.Write(ctx, room.Users[0].Ws, envelope{"message": username + " joined room!"})
		if err != nil {
			app.logger.Error(err.Error())
		}
		err = wsjson.Write(ctx, conn, envelope{"message": "Room " + name + "successfully joined!"})
		if err != nil {
			app.logger.Error(err.Error())
		}
	} else {
		return "", errors.New("Room is full")
	}

	app.rooms[name] = room
	return username + " successfully joined room!", nil
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
					app.logger.Info("found him")
					name = roomName
					username = user.Name
				}
			}
		}
	}

	room, ok := app.rooms[name]
	if !ok {
		return errors.New("Room does not exist (leave)")
	}

	if username == room.Users[1].Name {
		room.Users[1] = User{
			Name: "",
		}
		err := wsjson.Write(ctx, room.Users[0].Ws, envelope{"message": username + " left room."})
		if err != nil {
			app.logger.Error(err.Error())
		}
		return nil
	} else if username == room.Users[0].Name {
		if room.Users[1].Name != "" {
			err := wsjson.Write(ctx, room.Users[1].Ws, envelope{"message": room.Users[0].Name + " left room. Closing room and return to list..."})
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
