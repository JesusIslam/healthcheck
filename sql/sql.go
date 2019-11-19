package sql

import (
	"context"
	db "database/sql"
)

type SQLCheck struct {
	DB *db.DB
}

func New(db *db.DB) *SQLCheck {
	m := &SQLCheck{
		DB: db,
	}

	return m
}

func (m *SQLCheck) Check(ctx context.Context) (err error) {
	err = m.DB.PingContext(ctx)

	return
}
