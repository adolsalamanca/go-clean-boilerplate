package entities

import "time"

type Item struct {
	Id        string
	Name      string
	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
