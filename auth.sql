-- name: add-user
INSERT INTO auth (email, password)
  VALUES ($1, crypt($2, gen_salt('bf')))
;
