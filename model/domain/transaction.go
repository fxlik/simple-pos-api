package domain

import "time"

type Transaction struct {
	Id                  int32
	Code                string
	TransactionTypeId   int32
	SupplierId          int32
	TransactionStatusId int32
	CreatedBy           int32
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
