package models

import (
	"aikido/db"
	"aikido/forms"
	"context"
	"errors"

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

func (m GroupModel) Create(ctx context.Context, userID int64, form forms.CreateGroupForm) (int64, error) {
	if !userGroupModel.Exists(ctx, userID, 1) {
		return 0, errors.New("you can't create groups")
	}

	group := &Group{
		Name: form.Name,
	}
	err := db.GetDB().NewInsert().
		Model(group).
		Scan(ctx)
	return group.ID, err
}

func (m GroupModel) Delete(ctx context.Context, userID, id int64) error {
	if !userGroupModel.Exists(ctx, userID, 1) {
		return errors.New("you can't delete groups")
	}
	if id <= 2 {
		return errors.New("default groups can't be deleted")
	}

	tx, err := db.GetDB().Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// delete all userGroups
	_, err = tx.NewDelete().
		Model((*UserGroup)(nil)).
		Where("group_id = ?", id).
		Exec(ctx)
	if err != nil {
		return err
	}

	// delete group
	_, err = tx.NewDelete().
		Model((*Group)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (m GroupModel) Update(ctx context.Context, userID, id int64, form forms.UpdateGroupForm) error {
	if !userGroupModel.Exists(ctx, userID, 1) {
		return errors.New("you can't update groups")
	}
	if id <= 2 {
		return errors.New("default groups can't be updated")
	}

	_, err := db.GetDB().NewUpdate().
		Model((*Group)(nil)).
		Where("id = ?", id).
		Set("name = ?", form.Name).
		Exec(ctx)
	return err
}
