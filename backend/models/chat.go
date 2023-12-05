package models

import (
	"aikido/db"
	"aikido/forms"
	"context"
	"errors"
	"time"

	"github.com/uptrace/bun"
)

type Message struct {
	bun.BaseModel `bun:"table:messages,alias:m"`

	ID        int64     `json:"id" bun:",pk,autoincrement"`
	UserID    int64     `json:"-"`
	GroupID   int64     `json:"-"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`

	User *User `json:"user" bun:"rel:belongs-to,join:user_id=id"`
}

type ChatModel struct{}

func (m ChatModel) WriteMessage(ctx context.Context, userID, groupID int64, form forms.WriteMessageForm) error {
	// check if group exists
	groupCheck, err := db.GetDB().NewSelect().
		Model((*Group)(nil)).
		Where("id = ?", groupID).
		Exists(ctx)
	if err != nil {
		return err
	}
	if !groupCheck {
		return errors.New("group doesn't exist")
	}

	if !userGroupModel.Exists(ctx, userID, groupID) {
		return errors.New("you can't write to this chat")
	}

	// insert message
	message := &Message{
		UserID:    userID,
		GroupID:   groupID,
		Text:      form.Text,
		CreatedAt: time.Now(),
	}
	err = db.GetDB().NewInsert().
		Model(message).
		Scan(ctx)
	// TODO: send to websockets
	return err
}

func (m ChatModel) ReadMessages(ctx context.Context, userID, groupID int64, count, offset int) ([]*Message, error) {
	if !userGroupModel.Exists(ctx, userID, groupID) {
		return nil, errors.New("you can't read this chat")
	}

	messages := []*Message{}
	return messages, db.GetDB().NewSelect().
		Model(&messages).
		Where("group_id = ?", groupID).
		Relation("User").
		Limit(count).
		Offset(offset).
		Order("created_at DESC").
		Scan(ctx)

}
