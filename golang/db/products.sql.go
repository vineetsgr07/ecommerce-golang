// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: products.sql

package db

import (
	"context"
	"database/sql"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO product.products (name, description, status, image_url) 
VALUES ($1, $2, $3, $4) 
RETURNING id, name, description, status, image_url, created_at, updated_at, user_id
`

type CreateProductParams struct {
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Status      ProductProductStatus `json:"status"`
	ImageUrl    sql.NullString       `json:"image_url"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (ProductProduct, error) {
	row := q.db.QueryRow(ctx, createProduct,
		arg.Name,
		arg.Description,
		arg.Status,
		arg.ImageUrl,
	)
	var i ProductProduct
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Status,
		&i.ImageUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
	)
	return i, err
}

const deleteProductByID = `-- name: DeleteProductByID :exec
DELETE FROM product.products 
WHERE id = $1
`

func (q *Queries) DeleteProductByID(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteProductByID, id)
	return err
}

const findAllProducts = `-- name: FindAllProducts :exec
SELECT id, name, description, status, image_url, created_at, updated_at, user_id FROM product.products 
ORDER BY id ASC
`

func (q *Queries) FindAllProducts(ctx context.Context) error {
	_, err := q.db.Exec(ctx, findAllProducts)
	return err
}

const findProductByID = `-- name: FindProductByID :one
SELECT id, name, description, status, image_url, created_at, updated_at, user_id FROM product.products 
WHERE id = $1 
LIMIT 1
`

func (q *Queries) FindProductByID(ctx context.Context, id int64) (ProductProduct, error) {
	row := q.db.QueryRow(ctx, findProductByID, id)
	var i ProductProduct
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Status,
		&i.ImageUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
	)
	return i, err
}

const findProductsByStatus = `-- name: FindProductsByStatus :many
SELECT id, name, description, status, image_url, created_at, updated_at, user_id FROM product.products 
WHERE status = $1 
ORDER BY id DESC
`

func (q *Queries) FindProductsByStatus(ctx context.Context, status ProductProductStatus) ([]ProductProduct, error) {
	rows, err := q.db.Query(ctx, findProductsByStatus, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ProductProduct{}
	for rows.Next() {
		var i ProductProduct
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Status,
			&i.ImageUrl,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProduct = `-- name: UpdateProduct :one
UPDATE product.products 
SET name = $2, description = $3, status = $4, image_url = $5, updated_at = NOW() 
WHERE id = $1 
RETURNING id, name, description, status, image_url, created_at, updated_at, user_id
`

type UpdateProductParams struct {
	ID          int64                `json:"id"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Status      ProductProductStatus `json:"status"`
	ImageUrl    sql.NullString       `json:"image_url"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (ProductProduct, error) {
	row := q.db.QueryRow(ctx, updateProduct,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.Status,
		arg.ImageUrl,
	)
	var i ProductProduct
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Status,
		&i.ImageUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
	)
	return i, err
}
