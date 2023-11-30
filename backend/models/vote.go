package models

import (
	"aikido/db"
	"aikido/forms"
	"context"
	"errors"

	"github.com/uptrace/bun"
)

type Vote struct {
	bun.BaseModel `bun:"table:votes,alias:v"`

	UserID   int64 `bun:",pk"`
	PoolID   int64 `bun:",pk"`
	OptionID int64
}

type VoteModel struct{}

func (m VoteModel) Create(ctx context.Context, userID, poolID int64, form forms.CreateVoteForm) error {
	// check if already voted
	checkVote, err := db.GetDB().NewSelect().
		Model((*Vote)(nil)).
		Where("user_id = ? AND pool_id = ?", userID, poolID).
		Limit(1).
		Count(ctx)
	if err != nil {
		return err
	}
	if checkVote > 0 {
		return errors.New("already voted")
	}

	// check if option is related to pool
	checkOption, err := db.GetDB().NewSelect().
		Model((*Option)(nil)).
		Where("id = ? AND pool_id = ?", form.ID, poolID).
		Limit(1).
		Count(ctx)
	if err != nil {
		return err
	}
	if checkOption == 0 {
		return errors.New("invalid option")
	}

	// insert vote
	vote := &Vote{
		UserID:   userID,
		PoolID:   poolID,
		OptionID: form.ID,
	}
	_, err = db.GetDB().NewInsert().
		Model(vote).
		Exec(ctx)
	return err
}

func (m VoteModel) Delete(ctx context.Context, userID, poolID int64) error {
	_, err := db.GetDB().NewDelete().
		Model((*Vote)(nil)).
		Where("user_id = ? AND pool_id = ?", userID, poolID).
		Exec(ctx)
	return err
}
