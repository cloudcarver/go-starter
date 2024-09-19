// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: orgs.sql

package querier

import (
	"context"

	"github.com/google/uuid"
)

const createOrg = `-- name: CreateOrg :one
INSERT INTO orgs (
    name
) VALUES ($1) RETURNING id, name, owner_id, created_at, updated_at, deleted_at
`

func (q *Queries) CreateOrg(ctx context.Context, name string) (*Org, error) {
	row := q.db.QueryRow(ctx, createOrg, name)
	var i Org
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.OwnerID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const updateOrgOwnerID = `-- name: UpdateOrgOwnerID :exec
UPDATE orgs SET owner_id = $1 WHERE id = $2
`

type UpdateOrgOwnerIDParams struct {
	OwnerID uuid.NullUUID
	ID      uuid.UUID
}

func (q *Queries) UpdateOrgOwnerID(ctx context.Context, arg UpdateOrgOwnerIDParams) error {
	_, err := q.db.Exec(ctx, updateOrgOwnerID, arg.OwnerID, arg.ID)
	return err
}