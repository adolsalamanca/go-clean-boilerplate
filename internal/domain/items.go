package domain

import "time"

type Item struct {
	Id        int
	Name      string
	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
