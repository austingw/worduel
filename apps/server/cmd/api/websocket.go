package main

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
)

type Message struct {
	Content  string `json:"content"`
	Room     string `json:"room"`
	Type     string `json:"type"`
	Username string `json:"username"`
}

func (app *application) websocketHandler(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		OriginPatterns: []string{"localhost:5173"},
	})
	if err != nil {
		app.logger.Error(err.Error())
	}
	defer c.CloseNow()

	// Set the context as needed. Use of r.Context() is not recommended
	// to avoid surprising behavior (see http.Hijacker).
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	for {
		var input Message
		err = wsjson.Read(ctx, c, &input)

		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			err := app.leaveRoom(input.Room, input.Username, c)
			if err != nil {
				app.logger.Error(err.Error())
			}
			return
		}
		if websocket.CloseStatus(err) == websocket.StatusNormalClosure ||
			websocket.CloseStatus(err) == websocket.StatusGoingAway {
			err := app.leaveRoom(input.Room, input.Username, c)
			if err != nil {
				app.logger.Error(err.Error())
			}
			return
		}
		if err != nil {
			app.logger.Error(err.Error())
		}

		switch {
		case input.Type == "join":
			{
				msg, err := app.joinRoom(input.Room, input.Username, c)
				if err != nil {
					app.logger.Error(err.Error())
				}
				app.logger.Info(msg)
				err = wsjson.Write(ctx, c, envelope{"message": msg})
				if err != nil {
					app.logger.Error(err.Error())
				}
			}
		case input.Type == "attempt":
			{
				app.logger.Info("an attempt was made", input.Content)
			}
		case input.Type == "leave":
			{
				err := app.leaveRoom(input.Room, input.Username, c)
				if err != nil {
					app.logger.Error(err.Error())
				}
			}
		}

		// err = wsjson.Write(ctx, c, input)
		// if err != nil {
		// 	app.logger.Error(err.Error())
		// }
	}
}

func (app *application) sendJsonToPlayers(ctx context.Context, name, msg string) error {
	room, ok := app.rooms[name]
	if !ok {
		return errors.New("Room does not exist")
	}
	for _, user := range room.Users {
		err := wsjson.Write(ctx, user.Ws, envelope{"message": msg})
		if err != nil {
			return err
		}
	}
	return nil
}
