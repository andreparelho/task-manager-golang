package model

import "time"

type Task struct {
	Id          int
	Title       string
	Description string
	Status      string
	CreatedAt   time.Time
}
