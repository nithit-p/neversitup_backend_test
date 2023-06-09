CREATE TABLE IF NOT EXISTS auths (
  auth_id SERIAL PRIMARY KEY,
  user_id int NOT NULL,
  username text NOT NULL UNIQUE,
  email text NOT NULL UNIQUE,
  password_hash text NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  username text NOT NULL UNIQUE,
  email text NOT NULL UNIQUE,
  first_name text,
  last_name text
);

CREATE TABLE IF NOT EXISTS products (
  product_id SERIAL PRIMARY KEY,
  name text NOT NULL UNIQUE,
  description text,
  price int ,
  created_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS orders (
  order_id  SERIAL PRIMARY KEY,
  user_id  int NOT NULL,
  total_amount int,
  order_date TIMESTAMP,
  status text  NOT NULL

);

CREATE TABLE IF NOT EXISTS order_Items (
  order_item_id SERIAL PRIMARY KEY,
  order_id  int NOT NULL,
  product_id  int NOT NULL,
  quantity int,
  price int
);