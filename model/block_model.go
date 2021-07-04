package model

import "time"

type BlockModel struct {
	Id             int64
	Text           string
	Rate           int
	CreateDatetime time.Time
	tableName      struct{} `pg:"block"`
}
