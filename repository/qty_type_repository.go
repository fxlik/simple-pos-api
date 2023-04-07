package repository

import (
	"context"
	"database/sql"
	"fxlik/simple-post-api/model/domain"
)

type QtyTypeRepository interface {
	Save(ctx context.Context, tx *sql.Tx, qtyType domain.QtyType) domain.QtyType
	Update(ctx context.Context, tx *sql.Tx, qtyType domain.QtyType) domain.QtyType
	Delete(ctx context.Context, tx *sql.Tx, qtyType domain.QtyType)
	FindById(ctx context.Context, tx *sql.Tx, qtyTypeId int32) (domain.QtyType, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.QtyType
}
