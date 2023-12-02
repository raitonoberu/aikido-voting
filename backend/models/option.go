package models

import "github.com/uptrace/bun"

type Option struct {
	bun.BaseModel `bun:"table:options,alias:o"`

	ID     int64  `json:"id" bun:",pk,autoincrement"`
	PoolID int64  `json:"-"`
	Text   string `json:"text"`

	Count int `json:"count" bun:"-"`
}
