package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/ws", app.websocketHandler)
	mux.HandleFunc("GET /rooms", app.listRoomsHandler)
	mux.HandleFunc("POST /rooms", app.createRoomHandler)
	mux.HandleFunc("PUT /rooms", app.joinRoomHandler)

	return mux
}
