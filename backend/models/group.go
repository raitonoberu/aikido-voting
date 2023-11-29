package models

import (
	"aikido/db"
	"context"

	"github.com/uptrace/bun"
)

type Group struct {
	bun.BaseModel `bun:"table:groups,alias:g"`

	ID   int64  `json:"id" bun:",pk,autoincrement"`
	Name string `json:"name"`

	Users []*User `json:"users" bun:"m2m:user_groups,join:Group=User"`
}

type GroupModel struct{}

func (m GroupModel) All(ctx context.Context) ([]*Group, error) {
	var groups []*Group
	err := db.GetDB().NewSelect().
		Model(&groups).
		Scan(ctx)
	return groups, err
}
