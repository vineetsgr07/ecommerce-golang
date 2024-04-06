CREATE SCHEMA product;

CREATE TYPE product.product_status AS ENUM (
  'draft',
  'available',
  'out_of_stock',
  'discontinued'
);

CREATE TABLE product.products (
  id bigserial PRIMARY KEY,
  name varchar(255) NOT NULL,
  description text NOT NULL,
  status product.product_status NOT NULL,
  image_url varchar(255), -- Field to store the URL of the image associated with the product
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);