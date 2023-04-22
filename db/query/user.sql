-- name: CreateUser :one
INSERT INTO users (
    username,
    hashed_password,
    full_name,
    email,
    type,
    account_no,
    balance,
    currency
) VALUES (
    $1,$2,$3,$4,$5,$6,$7,$8
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1;

-- name: GetUserByAccNo :one
SELECT * FROM users
WHERE account_no = $1;

-- name: UpdateUser :one
UPDATE users 
SET
    hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
    password_changed_at = COALESCE(sqlc.narg(password_changed_at), password_changed_at),
    full_name = COALESCE(sqlc.narg(full_name), full_name),
    email = COALESCE(sqlc.narg(email), email)
WHERE
    username = sqlc.arg(username)
RETURNING *;
