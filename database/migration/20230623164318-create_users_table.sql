-- +migrate Up
CREATE TABLE users (
  id uuid UNIQUE PRIMARY KEY,
  sure_name varchar(255),
  full_name varchar(255),
  PASSWORD varchar(255),
  email varchar(255) UNIQUE,
  google_id varchar(50) UNIQUE,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP
);

-- +migrate Down
DROP TABLE users;