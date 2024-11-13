package repository

import (
	"errors"
	"sync"
	"time"

	"day_11/internal/model"
)

type Repo struct {
	events map[string]model.Event
	mu     sync.RWMutex
}

func New() *Repo {
	return &Repo{events: make(map[string]model.Event)}
}

func (r *Repo) CreateEvent(event model.Event) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.events[event.ID] = event
	return nil
}

func (r *Repo) UpdateEvent(event model.Event) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.events[event.ID]; !exists {
		return errors.New("event not found")
	}
	r.events[event.ID] = event
	return nil
}

func (r *Repo) DeleteEvent(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.events[id]; !exists {
		return errors.New("event not found")
	}
	delete(r.events, id)
	return nil
}

func (r *Repo) EventsForDay(date time.Time) ([]model.Event, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var events []model.Event
	for _, event := range r.events {
		if sameDay(event.Date, date) {
			events = append(events, event)
		}
	}
	return events, nil
}

func (r *Repo) EventsForWeek(startDate time.Time) ([]model.Event, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var events []model.Event
	for _, event := range r.events {
		if inSameWeek(event.Date, startDate) {
			events = append(events, event)
		}
	}
	return events, nil
}

func (r *Repo) EventsForMonth(startDate time.Time) ([]model.Event, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var events []model.Event
	for _, event := range r.events {
		if inSameMonth(event.Date, startDate) {
			events = append(events, event)
		}
	}
	return events, nil
}

func sameDay(d1, d2 time.Time) bool {
	return d1.Year() == d2.Year() && d1.YearDay() == d2.YearDay()
}

func inSameWeek(d1, start time.Time) bool {
	y1, w1 := d1.ISOWeek()
	y2, w2 := start.ISOWeek()
	return y1 == y2 && w1 == w2
}

func inSameMonth(d1, start time.Time) bool {
	return d1.Year() == start.Year() && d1.Month() == start.Month()
}
