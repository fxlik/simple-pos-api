package repository

import (
	"context"
	"database/sql"
	"fxlik/simple-post-api/model/domain"
)

type TransactionTypeRepository interface {
	Save(ctx context.Context, tx *sql.Tx, transactionType domain.TransactionType) domain.TransactionType
	Update(ctx context.Context, tx *sql.Tx, transactionType domain.TransactionType) domain.TransactionType
	Delete(ctx context.Context, tx *sql.Tx, transactionType domain.TransactionType)
	FindById(ctx context.Context, tx *sql.Tx, transactionTypeId int32) (domain.TransactionType, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.TransactionType
}
