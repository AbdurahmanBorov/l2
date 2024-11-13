package api

import (
	"encoding/json"
	"net/http"
	"time"

	"day_11/internal/model"
)

type service interface {
	CreateEvent(event model.Event) error
	UpdateEvent(event model.Event) error
	DeleteEvent(id string) error
	EventsForDay(date time.Time) ([]model.Event, error)
	EventsForWeek(startDate time.Time) ([]model.Event, error)
	EventsForMonth(startDate time.Time) ([]model.Event, error)
}

type API struct {
	service service
}

func New(service service) *API {
	return &API{service: service}
}

func (a *API) CreateEvent(w http.ResponseWriter, r *http.Request) {
	event, err := parseEvent(r)
	if err != nil {
		http.Error(w, `{"error": "invalid input"}`, http.StatusBadRequest)
		return
	}
	if err := a.service.CreateEvent(event); err != nil {
		http.Error(w, `{"error": "service error"}`, http.StatusServiceUnavailable)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"result": "event created"})
}

func (a *API) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	event, err := parseEvent(r)
	if err != nil {
		http.Error(w, `{"error": "invalid input"}`, http.StatusBadRequest)
		return
	}
	if err := a.service.UpdateEvent(event); err != nil {
		http.Error(w, `{"error": "service error"}`, http.StatusServiceUnavailable)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"result": "event updated"})
}

func (a *API) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, `{"error": "invalid input"}`, http.StatusBadRequest)
		return
	}
	if err := a.service.DeleteEvent(id); err != nil {
		http.Error(w, `{"error": "service error"}`, http.StatusServiceUnavailable)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"result": "event deleted"})
}

func (a *API) EventsForDay(w http.ResponseWriter, r *http.Request) {
	date, err := parseDate(r)
	if err != nil {
		http.Error(w, `{"error": "invalid date format"}`, http.StatusBadRequest)
		return
	}
	events, err := a.service.EventsForDay(date)
	if err != nil {
		http.Error(w, `{"error": "service error"}`, http.StatusServiceUnavailable)
		return
	}
	json.NewEncoder(w).Encode(events)
}

func (a *API) EventsForWeek(w http.ResponseWriter, r *http.Request) {
	date, err := parseDate(r)
	if err != nil {
		http.Error(w, `{"error": "invalid date format"}`, http.StatusBadRequest)
		return
	}
	events, err := a.service.EventsForWeek(date)
	if err != nil {
		http.Error(w, `{"error": "service error"}`, http.StatusServiceUnavailable)
		return
	}
	json.NewEncoder(w).Encode(events)
}

func (a *API) EventsForMonth(w http.ResponseWriter, r *http.Request) {
	date, err := parseDate(r)
	if err != nil {
		http.Error(w, `{"error": "invalid date format"}`, http.StatusBadRequest)
		return
	}
	events, err := a.service.EventsForMonth(date)
	if err != nil {
		http.Error(w, `{"error": "service error"}`, http.StatusServiceUnavailable)
		return
	}
	json.NewEncoder(w).Encode(events)
}

func parseEvent(r *http.Request) (model.Event, error) {
	r.ParseForm()
	id := r.FormValue("id")
	userID := r.FormValue("user_id")
	dateStr := r.FormValue("date")
	details := r.FormValue("details")

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return model.Event{}, err
	}

	return model.Event{ID: id, UserID: userID, Date: date, Details: details}, nil
}

func parseDate(r *http.Request) (time.Time, error) {
	dateStr := r.URL.Query().Get("date")
	return time.Parse("2006-01-02", dateStr)
}
