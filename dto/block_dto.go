package dto

import "time"

type BLockDto struct {
	Id             int64     `json:"id"`
	Text           string    `json:"text"`
	Rate           int       `json:"rate"`
	CreateDatetime time.Time `json:"createDatetime"`
}
