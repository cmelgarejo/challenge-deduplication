// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query_events.sql

package database

import (
	"context"
)

const createEvent = `-- name: CreateEvent :one
INSERT INTO "challenge"."events" (
    "hash",
    "block",
    "timestamp",
    "amount",
    "from"
  )
VALUES ($1, $2, $3, $4, $5)
RETURNING id, hash, block, timestamp, amount, "from"
`

type CreateEventParams struct {
	Hash      string `db:"hash" json:"hash"`
	Block     int32  `db:"block" json:"block"`
	Timestamp int32  `db:"timestamp" json:"timestamp"`
	Amount    string `db:"amount" json:"amount"`
	From      string `db:"from" json:"from"`
}

func (q *Queries) CreateEvent(ctx context.Context, arg CreateEventParams) (ChallengeEvent, error) {
	row := q.db.QueryRow(ctx, createEvent,
		arg.Hash,
		arg.Block,
		arg.Timestamp,
		arg.Amount,
		arg.From,
	)
	var i ChallengeEvent
	err := row.Scan(
		&i.ID,
		&i.Hash,
		&i.Block,
		&i.Timestamp,
		&i.Amount,
		&i.From,
	)
	return i, err
}

const listEvents = `-- name: ListEvents :many
SELECT id, hash, block, timestamp, amount, "from"
FROM "challenge"."events"
ORDER BY "timestamp" DESC
`

func (q *Queries) ListEvents(ctx context.Context) ([]ChallengeEvent, error) {
	rows, err := q.db.Query(ctx, listEvents)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ChallengeEvent
	for rows.Next() {
		var i ChallengeEvent
		if err := rows.Scan(
			&i.ID,
			&i.Hash,
			&i.Block,
			&i.Timestamp,
			&i.Amount,
			&i.From,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}