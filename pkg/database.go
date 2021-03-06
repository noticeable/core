package v1

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

// DB represents a database connection. It wraps a sqlx.DB to provide convenience methods.
type DB struct {
	sqlx.DB
}

// NewDB creates a new DB using an existing sqlx.DB connection.
func NewDB(db *sqlx.DB) *DB {
	return &DB{
		*db,
	}
}

// Selectx performs a select query using a squirrel SelectBuilder as an argument.
//
// This is a convenience wrapper. Any errors from squirrel or sqlx are returned as is.
func (db *DB) Selectx(dest interface{}, builder sq.SelectBuilder) error {
	sql, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	return db.Select(dest, sql, args...)
}

// Getx performs a get query using a squirrel SelectBuilder as an argument.
//
// This is a convenience wrapper. Any errors from squirrel or sqlx are returned as is.
func (db *DB) Getx(dest interface{}, builder sq.SelectBuilder) error {
	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	return db.Get(dest, query, args...)
}
