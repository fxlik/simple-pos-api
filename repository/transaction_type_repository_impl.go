package repository

import (
	"context"
	"database/sql"
	"errors"
	"fxlik/simple-post-api/helper"
	"fxlik/simple-post-api/model/domain"
)

type TransactionTypeRepositoryImpl struct {
}

func NewTransactionTypeRepositoryImpl() *TransactionTypeRepositoryImpl {
	return &TransactionTypeRepositoryImpl{}
}

func (repository *TransactionTypeRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, transactionType domain.TransactionType) domain.TransactionType {
	SQL := "INSERT INTO transaction_type(name, disabled) VALUES (?, ?)"
	result, err := tx.ExecContext(ctx, SQL, transactionType.Name, transactionType.Disabled)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	transactionType.Id = int32(id)
	return transactionType
}

func (repository *TransactionTypeRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, transactionType domain.TransactionType) domain.TransactionType {
	SQL := "UPDATE transaction_type SET name = ?, disabled = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, transactionType.Name, transactionType.Disabled, transactionType.Id)
	helper.PanicIfError(err)
	return transactionType
}

func (repository *TransactionTypeRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, transactionType domain.TransactionType) {
	SQL := "DELETE FROM transaction_type WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, transactionType.Id)
	helper.PanicIfError(err)
}

func (repository *TransactionTypeRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, transactionTypeId int32) (domain.TransactionType, error) {
	SQL := "SELECT id, name, disabled FROM transaction_type WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, transactionTypeId)
	helper.PanicIfError(err)
	defer rows.Close()

	transactionType := domain.TransactionType{}

	if rows.Next() {
		err := rows.Scan(&transactionType.Id, &transactionType.Name, &transactionType.Disabled)
		helper.PanicIfError(err)
		return transactionType, nil
	} else {
		return transactionType, errors.New("transaction_type is not found")
	}
}

func (repository *TransactionTypeRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.TransactionType {
	SQL := "SELECT * FROM transaction_type"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var transactionTypes []domain.TransactionType
	for rows.Next() {
		transactionType := domain.TransactionType{}
		err := rows.Scan(&transactionType.Id, &transactionType.Name, &transactionType.Disabled)
		helper.PanicIfError(err)
		transactionTypes = append(transactionTypes, transactionType)
	}
	return transactionTypes
}
