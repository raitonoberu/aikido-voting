package models

import (
	"aikido/db"
	"aikido/forms"
	"context"
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type Pool struct {
	bun.BaseModel `bun:"table:pools,alias:p"`

	ID          int64     `json:"id" bun:",pk,autoincrement"`
	UserID      int64     `json:"-"`
	GroupID     int64     `json:"-"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsAnonymous bool      `json:"is_anonymous"`
	CreatedAt   time.Time `json:"created_at"`
	ExpiresdAt  time.Time `json:"expires_at"`

	Vote int64 `json:"vote" bun:"-"`

	User *User `json:"user" bun:"rel:belongs-to,join:user_id=id"`
	// Group   *Group    `json:"group" bun:"rel:belongs-to,join:group_id=id"`
	Options []*Option `json:"options" bun:"rel:has-many,join:id=pool_id"`
}

type PoolModel struct{}

var userGroupModel = new(UserGroupModel)

func (m PoolModel) Create(ctx context.Context, userID int64, form forms.CreatePoolForm) (int64, error) {
	tx, err := db.GetDB().BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	// insert pool first to get pool_id
	pool := &Pool{
		UserID:      userID,
		GroupID:     form.GroupID,
		Name:        form.Name,
		Description: form.Description,
		IsAnonymous: *form.IsAnonymous,
		CreatedAt:   time.Now(),
		ExpiresdAt:  time.Now().Add(time.Hour * time.Duration(form.OpenFor)),
	}
	_, err = tx.NewInsert().
		Model(pool).
		Exec(ctx)
	if err != nil {
		return 0, err
	}

	// insert options
	for _, opt := range form.Options {
		option := &Option{
			PoolID: pool.ID,
			Text:   opt,
		}
		_, err = tx.NewInsert().
			Model(option).
			Exec(ctx)
		if err != nil {
			return 0, err
		}
	}

	return pool.ID, tx.Commit()
}

func (m PoolModel) Get(ctx context.Context, userID, id int64) (*Pool, error) {
	pool := &Pool{}
	err := db.GetDB().NewSelect().
		Model(pool).
		Where("p.id = ?", id).
		Relation("User").
		Relation("Options").
		Limit(1).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	for _, option := range pool.Options {
		// temporary solution
		count, err := db.GetDB().NewSelect().
			Model((*Vote)(nil)).
			Where("pool_id = ? AND option_id = ?", id, option.ID).
			Count(ctx)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		option.Count = count
	}

	// temporary solution
	userVote := &Vote{}
	err = db.GetDB().NewSelect().
		Model(userVote).
		Where("pool_id = ? AND user_id = ?", id, userID).
		Limit(1).
		Scan(ctx)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	pool.Vote = userVote.OptionID

	return pool, err
}

func (m PoolModel) Available(ctx context.Context, userID int64) ([]*Pool, error) {
	// get user groups
	groups, err := userGroupModel.ByUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	groupIDs := make([]int64, len(groups))
	for i, group := range groups {
		groupIDs[i] = group.GroupID
	}

	// get pools
	var pool []*Pool
	err = db.GetDB().NewSelect().
		Model(&pool).
		Where("group_id in (?)", bun.In(groupIDs)).
		Relation("User").
		Relation("Options").
		Scan(ctx)
	return pool, err
}

func (m PoolModel) Delete(ctx context.Context, id int64) error {
	tx, err := db.GetDB().Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// delete options
	_, err = tx.NewDelete().
		Model((*Option)(nil)).
		Where("pool_id = ?", id).
		Exec(ctx)
	if err != nil {
		return err
	}

	// delete pool
	_, err = tx.NewDelete().
		Model((*Pool)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return err
	}

	return tx.Commit()
}
