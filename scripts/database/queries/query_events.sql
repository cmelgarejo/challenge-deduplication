-- name: ListEvents :many
SELECT *
FROM "challenge"."events"
ORDER BY "timestamp" DESC;

-- name: CreateEvent :one
INSERT INTO "challenge"."events" (
    "hash",
    "block",
    "timestamp",
    "amount",
    "from"
  )
VALUES ($1, $2, $3, $4, $5)
RETURNING *;