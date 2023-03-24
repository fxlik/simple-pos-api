package domain

import "time"

type TransactionItem struct {
	Id            int32
	TransactionId int32
	ProductId     int32
	Qty           int32
	Price         int32
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
