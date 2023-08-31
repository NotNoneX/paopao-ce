// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: contacts.sql

package dbr

import (
	"context"
)

const createContact = `-- name: CreateContact :exec

INSERT INTO p_contact (user_id, friend_id, status, created_on) VALUES ($1, $2, $3, $4)
`

type CreateContactParams struct {
	UserID    int64
	FriendID  int64
	Status    int16
	CreatedOn int64
}

// ------------------------------------------------------------------------------
// contact_manager sql dml
// ------------------------------------------------------------------------------
func (q *Queries) CreateContact(ctx context.Context, arg *CreateContactParams) error {
	_, err := q.db.Exec(ctx, createContact,
		arg.UserID,
		arg.FriendID,
		arg.Status,
		arg.CreatedOn,
	)
	return err
}
