package migrator

import (
	"context"
	"database/sql"
	"log"
)

type (
	MySqlV1 struct {
		db *sql.DB
	}
)

func NewMySqlV1(db *sql.DB) *MySqlV1 {
	return &MySqlV1{
		db: db,
	}
}

func (x *MySqlV1) Init(ctx context.Context) error {
	tx, err := x.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		p := recover()
		if p == nil {
			if err := tx.Commit(); err != nil {
				log.Fatal(err)
			}
		} else {
			if err := tx.Rollback(); err != nil {
				log.Fatal(err)
			}
		}
	}()

	panic("implement me")
}
