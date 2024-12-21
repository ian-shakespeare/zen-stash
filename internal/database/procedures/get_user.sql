CREATE OR REPLACE PROCEDURE get_user(eml VARCHAR(255))
LANGUAGE SQL
AS $$
  SELECT
    user_id,
    first_name,
    last_name,
    email,
    password_digest,
    created_at
  FROM users
  WHERE email = eml;
$$;
