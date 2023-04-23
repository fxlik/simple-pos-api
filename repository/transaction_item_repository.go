package repository

import (
	"context"
	"database/sql"
	"fxlik/simple-post-api/model/domain"
)

type TransactionItemRepository interface {
	Save(ctx context.Context, tx *sql.Tx, item domain.TransactionItem) domain.TransactionItem
	Update(ctx context.Context, tx *sql.Tx, item domain.TransactionItem) domain.TransactionItem
	Delete(ctx context.Context, tx *sql.Tx, item domain.TransactionItem)
	FindById(ctx context.Context, tx *sql.Tx, transactionItemId int32) (domain.TransactionItem, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.TransactionItem
}
