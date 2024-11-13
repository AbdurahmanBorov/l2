package service

import (
	"time"

	"day_11/internal/model"
)

type repo interface {
	CreateEvent(event model.Event) error
	UpdateEvent(event model.Event) error
	DeleteEvent(id string) error
	EventsForDay(date time.Time) ([]model.Event, error)
	EventsForWeek(startDate time.Time) ([]model.Event, error)
	EventsForMonth(startDate time.Time) ([]model.Event, error)
}

type Service struct {
	repo repo
}

func New(repo repo) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateEvent(event model.Event) error {
	return s.repo.CreateEvent(event)
}

func (s *Service) UpdateEvent(event model.Event) error {
	return s.repo.UpdateEvent(event)
}

func (s *Service) DeleteEvent(id string) error {
	return s.repo.DeleteEvent(id)
}

func (s *Service) EventsForDay(date time.Time) ([]model.Event, error) {
	return s.repo.EventsForDay(date)
}

func (s *Service) EventsForWeek(startDate time.Time) ([]model.Event, error) {
	return s.repo.EventsForWeek(startDate)
}

func (s *Service) EventsForMonth(startDate time.Time) ([]model.Event, error) {
	return s.repo.EventsForMonth(startDate)
}
