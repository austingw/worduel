package main

import (
	"context"
	"errors"
	"net/http"
	"time"
	"worduel/internal/data"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
)

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
		var input data.Message
		err = wsjson.Read(ctx, c, &input)

		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			return
		}
		if websocket.CloseStatus(err) == websocket.StatusNormalClosure ||
			websocket.CloseStatus(err) == websocket.StatusGoingAway {
			return
		}
		if err != nil {
			app.logger.Error(err.Error())
		}

		err = wsjson.Write(ctx, c, input)
		if err != nil {
			app.logger.Error(err.Error())
		}
	}
}
