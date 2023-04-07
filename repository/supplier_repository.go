package repository

import (
	"context"
	"database/sql"
	"fxlik/simple-post-api/model/domain"
)

type SupplierRepository interface {
	Save(ctx context.Context, tx *sql.Tx, supplier domain.Supplier) domain.Supplier
	Update(ctx context.Context, tx *sql.Tx, supplier domain.Supplier) domain.Supplier
	Delete(ctx context.Context, tx *sql.Tx, supplier domain.Supplier)
	FindById(ctx context.Context, tx *sql.Tx, supplierId int32) (domain.Supplier, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Supplier
}
