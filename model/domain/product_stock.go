package domain

import "time"

type ProductStock struct {
	Id            int32
	ProductId     int32
	Qty           int32
	TransactionId int32
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
