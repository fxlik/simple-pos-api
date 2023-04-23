package repository

import (
	"context"
	"database/sql"
	"errors"
	"fxlik/simple-post-api/helper"
	"fxlik/simple-post-api/model/domain"
)

type TransactionRepositoryImpl struct {
}

func NewTransactionRepositoryImpl() *TransactionRepositoryImpl {
	return &TransactionRepositoryImpl{}
}

func (repository *TransactionRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) domain.Transaction {
	SQL := "INSERT INTO transaction(code, transaction_type_id, supplier_id, transaction_status_id, created_by, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, transaction.Code, transaction.TransactionTypeId, transaction.SupplierId, transaction.TransactionStatusId, transaction.CreatedBy, transaction.CreatedAt, transaction.UpdatedAt)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	transaction.Id = int32(id)
	return transaction
}

func (repository *TransactionRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) domain.Transaction {
	SQL := "UPDATE transaction SET code = ?, transaction_type_id = ?, supplier_id = ?, transaction_status_id = ?, created_by = ?, updated_at = ? WHERE  id = ?"
	_, err := tx.ExecContext(ctx, SQL, transaction.Code, transaction.TransactionTypeId, transaction.SupplierId, transaction.TransactionStatusId,
		transaction.CreatedBy, transaction.UpdatedAt, transaction.Id)
	helper.PanicIfError(err)
	return transaction
}

func (repository *TransactionRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) {
	SQL := "DELETE FROM transaction WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, transaction.Id)
	helper.PanicIfError(err)
}

func (repository *TransactionRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, transactionId int32) (domain.Transaction, error) {
	SQL := "SELECT id,code, transaction_type_id, supplier_id, transaction_status_id, created_by, created_at, updated_at FROM transaction WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, transactionId)
	helper.PanicIfError(err)
	defer rows.Close()

	transaction := domain.Transaction{}
	if rows.Next() {
		err := rows.Scan(&transaction.Id, &transaction.Code, &transaction.TransactionTypeId, &transaction.SupplierId,
			&transaction.TransactionStatusId, &transaction.CreatedBy, &transaction.CreatedAt, &transaction.UpdatedAt)
		helper.PanicIfError(err)
		return transaction, nil
	} else {
		return transaction, errors.New("transaction is not found")
	}
}

func (repository *TransactionRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Transaction {
	SQL := "SELECT * FROM transaction"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var transactions []domain.Transaction
	for rows.Next() {
		transaction := domain.Transaction{}
		err := rows.Scan(&transaction.Id, &transaction.Code, &transaction.TransactionTypeId, &transaction.SupplierId,
			&transaction.TransactionStatusId, &transaction.CreatedBy, &transaction.CreatedAt, &transaction.UpdatedAt)
		helper.PanicIfError(err)
		transactions = append(transactions, transaction)
	}
	return transactions
}
