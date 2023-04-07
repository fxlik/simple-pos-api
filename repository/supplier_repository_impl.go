package repository

import (
	"context"
	"database/sql"
	"errors"
	"fxlik/simple-post-api/helper"
	"fxlik/simple-post-api/model/domain"
)

type SupplierRepositoryImpl struct {
}

func NewSupplierRepositoryImpl() *SupplierRepositoryImpl {
	return &SupplierRepositoryImpl{}
}

func (repository *SupplierRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, supplier domain.Supplier) domain.Supplier {
	SQL := "INSERT INTO suppliers(name, email, telp, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, supplier.Name, supplier.Email, supplier.Telp, supplier.CreatedAt, supplier.UpdatedAt)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	supplier.Id = int32(id)
	return supplier
}

func (repository *SupplierRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, supplier domain.Supplier) domain.Supplier {
	SQL := "UPDATE suppliers SET name = ?, email = ?, telp = ?, updated_at = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, supplier.Name, supplier.Email, supplier.Telp, supplier.UpdatedAt, supplier.Id)
	helper.PanicIfError(err)
	return supplier
}

func (repository *SupplierRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, supplier domain.Supplier) {
	SQL := "DELETE FROM suppliers WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, supplier.Id)
	helper.PanicIfError(err)
}

func (repository *SupplierRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, supplierId int32) (domain.Supplier, error) {
	SQL := "SELECT id, name, email, telp, created_at, updated_at from suppliers WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, supplierId)
	helper.PanicIfError(err)
	defer rows.Close()

	supplier := domain.Supplier{}
	if rows.Next() {
		err := rows.Scan(&supplier.Id, &supplier.Name, &supplier.Email, &supplier.Telp, &supplier.CreatedAt, &supplier.UpdatedAt)
		helper.PanicIfError(err)
		return supplier, nil
	} else {
		return supplier, errors.New("supplier is not found")
	}
}

func (repository *SupplierRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Supplier {
	SQL := "SELECT * FROM suppliers"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var suppliers []domain.Supplier
	for rows.Next() {
		supplier := domain.Supplier{}
		err := rows.Scan(&supplier.Id, &supplier.Name, &supplier.Email, &supplier.Telp, &supplier.CreatedAt, &supplier.UpdatedAt)
		helper.PanicIfError(err)
		suppliers = append(suppliers, supplier)
	}
	return suppliers
}
