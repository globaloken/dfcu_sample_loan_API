-- name: CreateLoan :one
INSERT INTO loans (
    username,
    amount
) VALUES (
    $1,$2
) RETURNING *;

-- name: GetLoanByAccUsername :many
SELECT * FROM loans
WHERE username = $1;
