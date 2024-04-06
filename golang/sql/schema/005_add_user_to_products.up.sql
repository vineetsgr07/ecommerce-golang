ALTER TABLE product.products
ADD COLUMN user_id bigint REFERENCES users(id);