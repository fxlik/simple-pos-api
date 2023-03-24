package domain

import "time"

type Category struct {
	Id        int32
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
