CREATE TABLE IF NOT EXISTS users  (
  id SERIAL PRIMARY KEY,
  username VARCHAR(20) UNIQUE,
  password VARCHAR(255),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);