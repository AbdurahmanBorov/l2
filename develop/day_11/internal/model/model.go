package model

import "time"

type Event struct {
	ID      string
	UserID  string
	Date    time.Time
	Details string
}
