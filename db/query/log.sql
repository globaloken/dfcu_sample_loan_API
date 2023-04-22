-- name: CreateLog :exec
INSERT INTO logs (
    username,
    type
) VALUES (
    $1,$2
);

-- name: GetLogs :many
SELECT * FROM logs;

-- name: GetLogStats :many
SELECT COUNT(*), type FROM logs GROUP BY type;
