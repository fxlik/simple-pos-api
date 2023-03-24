package domain

import "time"

type Invoice struct {
	Id            int32
	SubTotal      int32
	TransactionId int32
	CreatedBy     int32
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
