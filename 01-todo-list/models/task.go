package models

import "time"

type Task struct {
	ID          string    `csv:"ID"`
	Description string    `csv:"Description"`
	CreatedAt   time.Time `csv:"CreatedAt"`
	IsComplete  bool      `csv:"IsComplete"`
} 