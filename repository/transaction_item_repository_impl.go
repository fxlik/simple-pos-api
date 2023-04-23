package repository

import (
	"context"
	"database/sql"
	"errors"
	"fxlik/simple-post-api/helper"
	"fxlik/simple-post-api/model/domain"
)

type TransactionItemRepositoryImpl struct {
}

func NewTransactionItemRepositoryImpl() *TransactionItemRepositoryImpl {
	return &TransactionItemRepositoryImpl{}
}

func (repository *TransactionItemRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, item domain.TransactionItem) domain.TransactionItem {
	SQL := "INSERT INTO transaction_items(transaction_id, product_id, qty, price, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, item.TransactionId, item.ProductId, item.Qty, item.Price, item.CreatedAt, item.UpdatedAt)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	item.Id = int32(id)
	return item
}

func (repository *TransactionItemRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, item domain.TransactionItem) domain.TransactionItem {
	SQL := "UPDATE transaction_items SET transaction_id = ?, product_id = ?, qty = ?, price = ?, updated_at = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, item.TransactionId, item.ProductId, item.Qty, item.Price, item.UpdatedAt, item.Id)
	helper.PanicIfError(err)
	return item
}

func (repository *TransactionItemRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, item domain.TransactionItem) {
	SQL := "DELETE FROM transaction_items WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, item.Id)
	helper.PanicIfError(err)
}

func (repository *TransactionItemRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, transactionItemId int32) (domain.TransactionItem, error) {
	SQL := "SELECT id, transaction_id, product_id, qty, price, created_at, updated_at FROM transaction_items WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, transactionItemId)
	helper.PanicIfError(err)
	defer rows.Close()

	transactionItem := domain.TransactionItem{}
	if rows.Next() {
		err := rows.Scan(&transactionItem.Id, &transactionItem.TransactionId, &transactionItem.ProductId, &transactionItem.Qty,
			&transactionItem.Price, &transactionItem.CreatedAt, &transactionItem.UpdatedAt)
		helper.PanicIfError(err)
		return transactionItem, nil
	} else {
		return transactionItem, errors.New("transaction_item is not found")
	}
}

func (repository *TransactionItemRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.TransactionItem {
	SQL := "SELECT * FROM transaction_items"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var transactionItems []domain.TransactionItem
	for rows.Next() {
		transactionItem := domain.TransactionItem{}
		err := rows.Scan(&transactionItem.Id, &transactionItem.TransactionId, &transactionItem.ProductId, &transactionItem.Qty,
			&transactionItem.Price, &transactionItem.CreatedAt, &transactionItem.UpdatedAt)
		helper.PanicIfError(err)
		transactionItems = append(transactionItems, transactionItem)
	}
	return transactionItems
}
