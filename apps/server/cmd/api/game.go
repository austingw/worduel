package main

import (
	"context"
	"errors"
	"sync"
	"time"
	words "worduel/internal"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
)

func (app *application) checkAnswer(attempt Message, conn *websocket.Conn) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	room, ok := app.rooms[attempt.Room]
	if !ok {
		return errors.New("Room does not exist")
	}

	if attempt.Content != room.CurrentWord {
		for _, user := range room.Users {
			if user.Ws != conn {
				err := wsjson.Write(ctx, user.Ws, envelope{"message": attempt.Username + " just made a guess!"})
				if err != nil {
					app.logger.Error(err.Error())
				}
			}
		}
	}

	if attempt.Content == room.CurrentWord {
		var wg sync.WaitGroup
		for i, user := range room.Users {
			wg.Add(1)
			go func(user User, i int) {
				defer wg.Done()
				if user.Ws == conn {
					err := wsjson.Write(ctx, user.Ws, envelope{"type": "end", "message": "You won!"})
					if err != nil {
						app.logger.Error(err.Error())
					}
					room.Users[i].Score++
				} else {
					err := wsjson.Write(ctx, user.Ws, envelope{"type": "end", "message": attempt.Username + " won!"})
					if err != nil {
						app.logger.Error(err.Error())
					}
				}
			}(user, i)
		}
		wg.Wait()
		time.AfterFunc(5*time.Second, func() {
			app.writeJSONToRoom(room.Name, envelope{"type": "start", "message": "Next round will begin shortly..."})
			room.CurrentWord = words.NewWord()
		})
		app.rooms[room.Name] = room
	}
	return nil
}
