package domain

import "time"

type ProductPrice struct {
	Id        int32
	Price     int32
	CreatedBy int32
	CreatedAt time.Time
	UpdatedAt time.Time
}
