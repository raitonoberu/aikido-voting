package models

import (
	"aikido/db"
	"aikido/forms"
	"context"

	"github.com/uptrace/bun"
)

type Group struct {
	bun.BaseModel `bun:"table:groups,alias:g"`

	ID   int64  `json:"id" bun:",pk,autoincrement"`
	Name string `json:"name"`

	Users []*User `json:"-" bun:"m2m:user_groups,join:Group=User"`
}

type GroupModel struct{}

func (m GroupModel) All(ctx context.Context) ([]*Group, error) {
	groups := []*Group{}
	err := db.GetDB().NewSelect().
		Model(&groups).
		Scan(ctx)
	return groups, err
}

func (m GroupModel) Create(ctx context.Context, form forms.CreateGroupForm) (int64, error) {
	// TODO: limit to admin only
	group := &Group{
		Name: form.Name,
	}
	err := db.GetDB().NewInsert().
		Model(group).
		Scan(ctx)
	return group.ID, err
}

func (m GroupModel) Delete(ctx context.Context, id int64) error {
	// TODO: limit to admin only
	_, err := db.GetDB().NewDelete().
		Model((*Group)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}

func (m GroupModel) Update(ctx context.Context, id int64, form forms.UpdateGroupForm) error {
	// TODO: limit to admin only
	_, err := db.GetDB().NewUpdate().
		Model((*Group)(nil)).
		Where("id = ?", id).
		Set("name = ?", form.Name).
		Exec(ctx)
	return err
}
