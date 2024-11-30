CREATE OR REPLACE PROCEDURE create_user(fname VARCHAR(255), lname VARCHAR(255), eml VARCHAR(255), pwd VARCHAR)
LANGUAGE SQL
AS $$
  INSERT INTO users (first_name, last_name, email, password_digest)
  VALUES (fname, lname, eml, pwd);
$$;
