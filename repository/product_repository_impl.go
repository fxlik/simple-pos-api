package repository

import (
	"context"
	"database/sql"
	"errors"
	"fxlik/simple-post-api/helper"
	"fxlik/simple-post-api/model/domain"
)

type ProductRepositoryImpl struct {
}

func NewProductRepositoryImpl() *ProductRepositoryImpl {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "INSERT INTO products(category_id, name, code, image, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, product.CategoryId, product.Name, product.Code, product.Image, product.CreatedAt, product.UpdatedAt)
	helper.PanicIfError(err)

	id, errId := result.LastInsertId()
	helper.PanicIfError(errId)
	product.Id = int32(id)
	return product
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "UPDATE products SET category_id = ?, name = ?, code = ?, image = ?, updated_at = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.CategoryId, product.Name, product.Code, product.Image, product.UpdatedAt, product.Id)
	helper.PanicIfError(err)
	return product
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product domain.Product) {
	SQL := "DELETE FROM products WHERE id =?"
	_, err := tx.ExecContext(ctx, SQL, product.Id)
	helper.PanicIfError(err)
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int32) (domain.Product, error) {
	SQL := "SELECT id, category_id, name, code, image, created_at, updated_at FROM products WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, productId)
	helper.PanicIfError(err)
	defer rows.Close()

	product := domain.Product{}
	if rows.Next() {
		err := rows.Scan(&product.Id, &product.CategoryId, &product.Name, &product.Code, &product.Image, &product.CreatedAt, &product.UpdatedAt)
		helper.PanicIfError(err)
		return product, nil
	} else {
		return product, errors.New("product is not found")
	}
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Product {
	SQL := "SELECT * FROM products"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		product := domain.Product{}
		err := rows.Scan(&product.Id, &product.CategoryId, &product.Name, &product.Code, &product.Image, &product.CreatedAt, &product.UpdatedAt)
		helper.PanicIfError(err)

		products = append(products, product)
	}
	return products
}
