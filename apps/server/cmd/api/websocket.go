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
		OriginPatterns: []string{"localhost:5173", "worduel-dusky.vercel.app"},
	})
	if err != nil {
		app.logger.Error(err.Error())
	}
	defer c.CloseNow()

	// Set the context as needed. Use of r.Context() is not recommended
	// to avoid surprising behavior (see http.Hijacker).
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	for {
		var input Message
		err = wsjson.Read(ctx, c, &input)

		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			err := app.leaveRoom(input.Room, input.Username, c)
			if err != nil {
				app.logger.Error(err.Error(), "fie")
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
			err := app.leaveRoom(input.Room, input.Username, c)
			if err != nil {
				app.logger.Error(err.Error())
			}
			app.logger.Error(err.Error())
			return
		}

		switch {
		case input.Type == "join":
			{
				err := app.joinRoom(input.Room, input.Username, c)
				if err != nil {
					app.logger.Error(err.Error())
				}
			}
		case input.Type == "attempt":
			{
				err = app.checkAnswer(input, c)
				if err != nil {
					app.logger.Error(err.Error())
				}
			}
		case input.Type == "leave":
			{
				err := app.leaveRoom(input.Room, input.Username, c)
				if err != nil {
					app.logger.Error(err.Error())
				}
			}
		}
	}
}
