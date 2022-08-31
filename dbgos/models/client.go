package models

import (
	"context"
	"database/sql"
)

type (
	Client struct {
		Root *sql.DB
	}
)

func NewClient() error {
	db, err := sql.Open("mysql", "")
	if err != nil {
		return err
	}

	println(db.Stats())

	ctx := context.TODO()
	conn, err := db.Conn(ctx)
	if err != nil {
		return err
	}

	defer conn.Close()

	return nil
}
