package repository

import (
	"context"
	"database/sql"
	"errors"
	"fxlik/simple-post-api/helper"
	"fxlik/simple-post-api/model/domain"
)

type ProductPriceRepositoryImpl struct {
}

func NewProductPriceRepositoryImpl() *ProductPriceRepositoryImpl {
	return &ProductPriceRepositoryImpl{}
}

func (repository *ProductPriceRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, productPrice domain.ProductPrice) domain.ProductPrice {
	SQL := "INSERT INTO product_price(product_id, price, created_by, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, productPrice.ProductId, productPrice.Price, productPrice.CreatedBy, productPrice.CreatedAt, productPrice.UpdatedAt)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	productPrice.Id = int32(id)
	return productPrice
}

func (repository *ProductPriceRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, productPrice domain.ProductPrice) domain.ProductPrice {
	SQL := "UPDATE product_price SET product_id = ?, price = ?, created_by = ?, updated_at = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, productPrice.ProductId, productPrice.Price, productPrice.CreatedBy, productPrice.UpdatedAt, productPrice.Id)
	helper.PanicIfError(err)
	return productPrice
}

func (repository *ProductPriceRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, productPrice domain.ProductPrice) {
	SQL := "DELETE FROM product_price WHERE id =?"
	_, err := tx.ExecContext(ctx, SQL, productPrice.Id)
	helper.PanicIfError(err)
}

func (repository *ProductPriceRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productPriceId int32) (domain.ProductPrice, error) {
	SQL := "SELECT id, product_id, price, created_by, created_at, updated_at FROM product_price WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, productPriceId)
	helper.PanicIfError(err)
	defer rows.Close()

	productPrice := domain.ProductPrice{}
	if rows.Next() {
		err := rows.Scan(&productPrice.Id, &productPrice.ProductId, &productPrice.Price, &productPrice.CreatedBy, &productPrice.CreatedAt, &productPrice.UpdatedAt)
		helper.PanicIfError(err)
		return productPrice, nil
	} else {
		return productPrice, errors.New("product_price is not found")
	}
}

func (repository *ProductPriceRepositoryImpl) FindOneByProductId(ctx context.Context, tx *sql.Tx, productId int32) (domain.ProductPrice, error) {
	SQL := "SELECT id, product_id, price, created_by, created_at, updated_at FROM product_price WHERE product_id = ? ORDER BY created_at DESC LIMIT 1"
	rows, err := tx.QueryContext(ctx, SQL, productId)
	helper.PanicIfError(err)
	defer rows.Close()

	productPrice := domain.ProductPrice{}
	if rows.Next() {
		err := rows.Scan(&productPrice.Id, &productPrice.ProductId, &productPrice.Price, &productPrice.CreatedBy, &productPrice.CreatedAt, &productPrice.UpdatedAt)
		helper.PanicIfError(err)
		return productPrice, nil
	} else {
		return productPrice, errors.New("product_price is not found")
	}
}

func (repository *ProductPriceRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.ProductPrice {
	SQL := "SELECT * FROM product_price"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var productPrices []domain.ProductPrice
	for rows.Next() {
		productPrice := domain.ProductPrice{}
		err := rows.Scan(&productPrice.Id, &productPrice.ProductId, &productPrice.Price, &productPrice.CreatedBy, &productPrice.CreatedAt, &productPrice.UpdatedAt)
		helper.PanicIfError(err)

		productPrices = append(productPrices, productPrice)
	}
	return productPrices
}
