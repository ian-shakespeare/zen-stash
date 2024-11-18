CREATE TABLE IF NOT EXISTS users (
  user_id UUID PRIMARY KEY DEFAULT get_random_uuid()
);

CREATE TABLE IF NOT EXISTS files (
  file_id UUID PRIMARY KEY DEFAULT get_random_uuid(),
  name VARCHAR NOT NULL,
  mime VARCHAR NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT now,
  user_id UUID NOT NULL REFERENCES users(user_id)
);
