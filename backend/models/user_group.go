package models

import (
	"aikido/db"
	"context"

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

func (m UserGroupModel) Add(ctx context.Context, userID, groupID int64) error {
	_, err := db.GetDB().NewInsert().
		Model(&UserGroup{
			UserID:  userID,
			GroupID: groupID,
		}).
		Exec(ctx)
	return err
}

func (m UserGroupModel) Remove(ctx context.Context, userID, groupID int64) error {
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
