package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/coder/websocket"
	"github.com/rs/cors"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type User struct {
	Name  string          `json:"name"`
	Ws    *websocket.Conn `json:"-"`
	Score int             `json:"score"`
}

type Room struct {
	Name        string  `json:"name"`
	Users       [2]User `json:"users"`
	CurrentWord string  `json:"currentWord"`
}

type application struct {
	config config
	logger *slog.Logger
	rooms  map[string]Room
	mu     sync.RWMutex
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		config: cfg,
		logger: logger,
		rooms:  make(map[string]Room),
		mu:     sync.RWMutex{},
	}

	handler := cors.Default().Handler(app.routes())

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server", "addr", srv.Addr, "env", cfg.env)

	err := srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
