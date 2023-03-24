package domain

import "time"

type User struct {
	Id        int32
	LevelId   int32
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
