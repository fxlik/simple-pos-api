package repository

import (
	"context"
	"database/sql"
	"fxlik/simple-post-api/model/domain"
)

type ProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	Delete(ctx context.Context, tx *sql.Tx, product domain.Product)
	FindById(ctx context.Context, tx *sql.Tx, productId int32) (domain.Product, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Product
}
