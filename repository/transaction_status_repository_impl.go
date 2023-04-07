package repository

import (
	"context"
	"database/sql"
	"errors"
	"fxlik/simple-post-api/helper"
	"fxlik/simple-post-api/model/domain"
)

type TransactionStatusRepositoryImpl struct {
}

func NewTransactionStatusRepositoryImpl() *TransactionStatusRepositoryImpl {
	return &TransactionStatusRepositoryImpl{}
}

func (repository *TransactionStatusRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, transactionStatus domain.TransactionStatus) domain.TransactionStatus {
	SQL := "INSERT INTO transaction_status(name, disabled) VALUES (?, ?)"
	result, err := tx.ExecContext(ctx, SQL, transactionStatus.Name, transactionStatus.Disabled)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	transactionStatus.Id = int32(id)
	return transactionStatus
}

func (repository *TransactionStatusRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, transactionStatus domain.TransactionStatus) domain.TransactionStatus {
	SQL := "UPDATE transaction_status SET name = ?, disabled = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, transactionStatus.Name, transactionStatus.Disabled, transactionStatus.Id)
	helper.PanicIfError(err)
	return transactionStatus
}

func (repository *TransactionStatusRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, transactionStatus domain.TransactionStatus) {
	SQL := "DELETE FROM transaction_status WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, transactionStatus.Id)
	helper.PanicIfError(err)
}

func (repository *TransactionStatusRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, transactionStatusId int32) (domain.TransactionStatus, error) {
	SQL := "SELECT id, name, disabled FROM transaction_status WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, transactionStatusId)
	helper.PanicIfError(err)
	defer rows.Close()

	transactionStatus := domain.TransactionStatus{}

	if rows.Next() {
		err := rows.Scan(&transactionStatus.Id, &transactionStatus.Name, &transactionStatus.Disabled)
		helper.PanicIfError(err)
		return transactionStatus, nil
	} else {
		return transactionStatus, errors.New("transaction_status is not found")
	}
}

func (repository *TransactionStatusRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.TransactionStatus {
	SQL := "SELECT * FROM transaction_status"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var transactionStatuses []domain.TransactionStatus
	for rows.Next() {
		transactionStatus := domain.TransactionStatus{}
		err := rows.Scan(&transactionStatus.Id, &transactionStatus.Name, &transactionStatus.Disabled)
		helper.PanicIfError(err)
		transactionStatuses = append(transactionStatuses, transactionStatus)
	}
	return transactionStatuses
}
