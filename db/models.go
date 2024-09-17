// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Book struct {
	ID        int32
	Title     pgtype.Text
	Author    pgtype.Text
	Publisher pgtype.Text
	Price     pgtype.Int4
}

type SchemaMigration struct {
	Version int64
	Dirty   bool
}
