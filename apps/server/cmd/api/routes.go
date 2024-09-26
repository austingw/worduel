package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/ws", app.enableCORS(http.HandlerFunc(app.websocketHandler)))
	mux.Handle("GET /rooms", app.enableCORS(http.HandlerFunc(app.listRoomsHandler)))
	mux.HandleFunc("POST /rooms", app.createRoomHandler)
	mux.HandleFunc("PUT /rooms", app.joinRoomHandler)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})

	return mux
}
