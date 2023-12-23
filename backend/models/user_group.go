package models

import (
	"aikido/db"
	"aikido/forms"
	"context"
	"errors"

	"github.com/uptrace/bun"
)

type UserGroup struct {
	bun.BaseModel `bun:"table:user_groups,alias:ug"`

	UserID  int64 `bun:",pk"`
	GroupID int64 `bun:",pk"`

	User  *User  `bun:"rel:belongs-to,join:user_id=id"`
	Group *Group `bun:"rel:belongs-to,join:group_id=id"`
}

type UserGroupModel struct{}

func (m UserGroupModel) Add(ctx context.Context, userID, groupID int64, form forms.AddUserForm) error {
	if !m.Exists(ctx, userID, 1) {
		return errors.New("you can't add users to groups")
	}
	db := db.GetDB()

	// check if user exists
	userExists, err := db.NewSelect().
		Model((*User)(nil)).
		Where("id = ?", form.ID).
		Exists(ctx)
	if err != nil {
		return err
	}
	if !userExists {
		return errors.New("user doesn't exist")
	}

	// check if group exists
	groupExists, err := db.NewSelect().
		Model((*Group)(nil)).
		Where("id = ?", groupID).
		Exists(ctx)
	if err != nil {
		return err
	}
	if !groupExists {
		return errors.New("group doesn't exist")
	}

	// insert userGroup
	userGroup := &UserGroup{
		UserID:  form.ID,
		GroupID: groupID,
	}
	_, err = db.NewInsert().
		Model(userGroup).
		Exec(ctx)
	return err
}

func (m UserGroupModel) Remove(ctx context.Context, curUserID, userID, groupID int64) error {
	if !m.Exists(ctx, curUserID, 1) {
		return errors.New("you can't remove users from groups")
	}
	if groupID == 1 && userID == 1 {
		return errors.New("this admin can't be removed")
	}

	_, err := db.GetDB().NewDelete().
		Model((*UserGroup)(nil)).
		Where("user_id = ? AND group_id = ?", userID, groupID).
		Exec(ctx)
	return err
}

func (m UserGroupModel) ByUser(ctx context.Context, userID int64) ([]*UserGroup, error) {
	groups := []*UserGroup{}
	err := db.GetDB().NewSelect().
		Model(&groups).
		Where("user_id = ?", userID).
		Scan(ctx)
	return groups, err
}

func (m UserGroupModel) ByGroup(ctx context.Context, groupID int64) ([]*UserGroup, error) {
	groups := []*UserGroup{}
	err := db.GetDB().NewSelect().
		Model(&groups).
		Where("group_id = ?", groupID).
		Scan(ctx)
	return groups, err
}

func (m UserGroupModel) Exists(ctx context.Context, userID int64, groupID int64) bool {
	exists, _ := db.GetDB().NewSelect().
		Model((*UserGroup)(nil)).
		Where("user_id = ? AND group_id = ?", userID, groupID).
		Exists(ctx)
	return exists
}
