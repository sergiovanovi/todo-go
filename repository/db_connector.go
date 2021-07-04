package repository

import (
	"github.com/go-pg/pg/v10"
)

func connect() *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     "127.0.0.1:5432",
		User:     "postgres",
		Password: "password",
		Database: "todo",
	})
	return db
}
