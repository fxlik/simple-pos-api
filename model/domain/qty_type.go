package domain

import "time"

type QtyType struct {
	Id        int32
	Name      string
	Disabled  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
