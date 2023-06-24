-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
  id uuid NOT NULL DEFAULT uuid_generate_v4(),
  email varchar(255) UNIQUE NOT NULL,
  sure_name varchar(255),
  full_name varchar(255),
  PASSWORD varchar(255),
  google_id varchar(50) UNIQUE DEFAULT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  PRIMARY KEY (id)
);

INSERT INTO
  users(email, sure_name, full_name, PASSWORD)
VALUES
  (
    'rizal.arfiyan.23@gmail.com',
    'Rizal',
    'Muhamad Rizal Arfiyan',
    '$2y$10$DHwowZw4DY7A9.cHFQ9c0eRyQ6B9Zsy.iZT7E/00MJKTrA1FFVQ4W'
  );

-- +migrate Down
DROP TABLE users;