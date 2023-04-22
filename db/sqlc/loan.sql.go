// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: loan.sql

package db

import (
	"context"
)

const createLoan = `-- name: CreateLoan :one
INSERT INTO loans (
    username,
    amount
) VALUES (
    $1,$2
) RETURNING id, username, amount, created_at
`

type CreateLoanParams struct {
	Username string `json:"username"`
	Amount   int64  `json:"amount"`
}

func (q *Queries) CreateLoan(ctx context.Context, arg CreateLoanParams) (Loan, error) {
	row := q.db.QueryRowContext(ctx, createLoan, arg.Username, arg.Amount)
	var i Loan
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getLoanByAccUsername = `-- name: GetLoanByAccUsername :many
SELECT id, username, amount, created_at FROM loans
WHERE username = $1
`

func (q *Queries) GetLoanByAccUsername(ctx context.Context, username string) ([]Loan, error) {
	rows, err := q.db.QueryContext(ctx, getLoanByAccUsername, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Loan{}
	for rows.Next() {
		var i Loan
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}