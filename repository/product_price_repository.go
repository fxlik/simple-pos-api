package repository

import (
	"context"
	"database/sql"
	"fxlik/simple-post-api/model/domain"
)

type ProductPriceRepository interface {
	Save(ctx context.Context, tx *sql.Tx, productPrice domain.ProductPrice) domain.ProductPrice
	Update(ctx context.Context, tx *sql.Tx, productPrice domain.ProductPrice) domain.ProductPrice
	Delete(ctx context.Context, tx *sql.Tx, productPrice domain.ProductPrice)
	FindById(ctx context.Context, tx *sql.Tx, productPriceId int32) (domain.ProductPrice, error)
	FindOneByProductId(ctx context.Context, tx *sql.Tx, productId int32) (domain.ProductPrice, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.ProductPrice
}
