package repository

import (
	"context"
	"database/sql"
	"errors"
	"fxlik/simple-post-api/helper"
	"fxlik/simple-post-api/model/domain"
)

type QtyTypeRepositoryImpl struct {
}

func NewQtyTypeRepositoryImpl() *QtyTypeRepositoryImpl {
	return &QtyTypeRepositoryImpl{}
}

func (repository *QtyTypeRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, qtyType domain.QtyType) domain.QtyType {
	SQL := "INSERT INTO qty_type(name, disabled, created_at, updated_at) VALUES (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, qtyType.Name, qtyType.Disabled, qtyType.CreatedAt, qtyType.UpdatedAt)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	qtyType.Id = int32(id)
	return qtyType
}

func (repository *QtyTypeRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, qtyType domain.QtyType) domain.QtyType {
	SQL := "UPDATE qty_type SET name = ?, disabled = ?, updated_at = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, qtyType.Name, qtyType.Disabled, qtyType.UpdatedAt, qtyType.Id)
	helper.PanicIfError(err)
	return qtyType
}

func (repository *QtyTypeRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, qtyType domain.QtyType) {
	SQL := "DELETE FROM qty_type WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, qtyType.Id)
	helper.PanicIfError(err)
}

func (repository *QtyTypeRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, qtyTypeId int32) (domain.QtyType, error) {
	SQL := "SELECT id, name, disabled, created_at, updated_at from qty_type WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, qtyTypeId)
	helper.PanicIfError(err)
	defer rows.Close()

	qtyType := domain.QtyType{}
	if rows.Next() {
		err := rows.Scan(&qtyType.Id, &qtyType.Name, &qtyType.Disabled, &qtyType.CreatedAt, &qtyType.UpdatedAt)
		helper.PanicIfError(err)
		return qtyType, nil
	} else {
		return qtyType, errors.New("qty_type is not found")
	}
}

func (repository *QtyTypeRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.QtyType {
	SQL := "SELECT * FROM qty_type"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var qtyTypes []domain.QtyType
	for rows.Next() {
		qtyType := domain.QtyType{}
		err := rows.Scan(&qtyType.Id, &qtyType.Name, &qtyType.Disabled, &qtyType.CreatedAt, &qtyType.UpdatedAt)
		helper.PanicIfError(err)
		qtyTypes = append(qtyTypes, qtyType)
	}
	return qtyTypes
}
