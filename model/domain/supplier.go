package domain

import "time"

type Supplier struct {
	Id        int32
	Name      string
	Email     string
	Telp      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
