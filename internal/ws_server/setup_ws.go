package ws_server

import (
	"github.com/rs/zerolog/log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/thelemonday/smart-parking-iot-server/db"
)

func RunBackendServer(db db.DatabaseRepo) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	wsServer := SetupWebsocketServer(db)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		_, err := w.Write([]byte("Hello world!"))
		if err != nil {
			log.Error().Err(err).Msg("")
		}
	})

	r.Get("/ws", func(w http.ResponseWriter, r *http.Request) {
		wsServer.handleNewWsConnections(w, r)
	})

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal().Err(err).Msg("")
		return
	}
}
