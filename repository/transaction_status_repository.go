package repository

import (
	"context"
	"database/sql"
	"fxlik/simple-post-api/model/domain"
)

type TransactionStatusRepository interface {
	Save(ctx context.Context, tx *sql.Tx, transactionStatus domain.TransactionStatus) domain.TransactionStatus
	Update(ctx context.Context, tx *sql.Tx, transactionStatus domain.TransactionStatus) domain.TransactionStatus
	Delete(ctx context.Context, tx *sql.Tx, transactionStatus domain.TransactionStatus)
	FindById(ctx context.Context, tx *sql.Tx, transactionStatusId int32) (domain.TransactionStatus, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.TransactionStatus
}
