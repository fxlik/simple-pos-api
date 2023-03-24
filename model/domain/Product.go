package domain

import "time"

type Product struct {
	Id         int32
	CategoryId int32
	Name       string
	Code       string
	Image      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
