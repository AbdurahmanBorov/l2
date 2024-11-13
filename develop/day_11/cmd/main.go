package main

import (
	"log"
	"net/http"

	handlers "day_11/internal/api"
	"day_11/internal/config"
	"day_11/internal/repo"
	"day_11/internal/service"
)

func main() {
	cfg := config.NewConfig()

	db := repository.New()
	s := service.New(db)
	handler := handlers.New(s)

	mux := http.NewServeMux()
	mux.HandleFunc("/event/create", handler.CreateEvent)
	mux.HandleFunc("/event/update", handler.UpdateEvent)
	mux.HandleFunc("/event/delete", handler.DeleteEvent)
	mux.HandleFunc("/event/day", handler.EventsForDay)
	mux.HandleFunc("/event/week", handler.EventsForWeek)
	mux.HandleFunc("/event/month", handler.EventsForMonth)

	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      handler.Logging(mux),
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	log.Printf("Starting server on %s", cfg.Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
