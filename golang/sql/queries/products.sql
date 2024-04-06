-- name: UpdateProduct :one
UPDATE product.products 
SET name = $2, description = $3, status = $4, image_url = $5, updated_at = NOW() 
WHERE id = $1 
RETURNING *;

-- name: FindProductsByStatus :many
SELECT * FROM product.products 
WHERE status = $1 
ORDER BY id DESC;

-- name: FindProductByID :one
SELECT * FROM product.products 
WHERE id = $1 
LIMIT 1;

-- name: DeleteProductByID :exec
DELETE FROM product.products 
WHERE id = $1;

-- name: FindAllProducts :exec
SELECT * FROM product.products 
ORDER BY id ASC;

-- name: CreateProduct :one
INSERT INTO product.products (name, description, status, image_url) 
VALUES ($1, $2, $3, $4) 
RETURNING *;