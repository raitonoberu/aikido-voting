package models

import (
	"aikido/db"
	"aikido/forms"
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID       int64  `json:"id" bun:",pk,autoincrement"`
	RoleID   int64  `json:"-"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`

	// Role   *Role    `json:"-" bun:"rel:belongs-to,join:role_id=id"`
	Groups []*Group `json:"-" bun:"m2m:user_groups,join:User=Group"`
}

type UserModel struct{}

var authModel = new(AuthModel)

func (m UserModel) Login(ctx context.Context, form forms.LoginForm) (*User, string, error) {
	user := &User{}
	err := db.GetDB().NewSelect().
		Model(user).
		Where("email = LOWER(?)", form.Email).
		Limit(1).
		Scan(ctx)
	if err != nil {
		return nil, "", err
	}

	// check password
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password), []byte(form.Password),
	)
	if err != nil {
		return nil, "", err
	}

	// create token
	token, err := authModel.CreateToken(user.ID)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (m UserModel) Register(ctx context.Context, form forms.RegisterForm) (*User, string, error) {
	db := db.GetDB()

	// check if user already exists
	checkUser, err := db.NewSelect().
		Model((*User)(nil)).
		Where("email = LOWER(?)", form.Email).
		Limit(1).
		Count(ctx)
	if err != nil {
		return nil, "", err
	}
	if checkUser > 0 {
		return nil, "", errors.New("email already exists")
	}

	// hash pasword
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(form.Password), bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, "", err
	}

	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, "", err
	}
	defer tx.Rollback()

	// insert user first to get user_id
	user := &User{
		RoleID:   1, // TODO
		Name:     form.Name,
		Email:    strings.ToLower(form.Email),
		Password: string(hashedPassword),
	}
	_, err = tx.NewInsert().
		Model(user).
		Exec(ctx)
	if err != nil {
		return nil, "", err
	}

	// insert default group
	defaultGroup := &UserGroup{
		UserID:  user.ID,
		GroupID: 1,
	}
	_, err = tx.NewInsert().
		Model(defaultGroup).
		Exec(ctx)
	if err != nil {
		return nil, "", err
	}

	if err := tx.Commit(); err != nil {
		return nil, "", err
	}

	// create token
	token, err := authModel.CreateToken(user.ID)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (m UserModel) Get(ctx context.Context, id int64) (*User, error) {
	user := &User{}
	err := db.GetDB().NewSelect().
		Model(user).
		Where("id = ?", id).
		Limit(1).
		Scan(ctx)
	return user, err
}

func (m UserModel) Update(ctx context.Context, userID int64, form forms.UpdateUserForm) error {
	query := db.GetDB().NewUpdate().
		Model((*User)(nil)).
		Where("id = ?", userID)

	if form.Password != "" {
		// TODO: logout all other sessions?
		hashedPassword, err := bcrypt.GenerateFromPassword(
			[]byte(form.Password), bcrypt.DefaultCost,
		)
		if err != nil {
			return err
		}
		query = query.Set("password = ?", string(hashedPassword))
	}
	if form.Email != "" {
		// check if email already exists
		checkUser, err := db.GetDB().NewSelect().
			Model((*User)(nil)).
			Where("email = LOWER(?)", form.Email).
			Count(ctx)
		if err != nil {
			return err
		}
		if checkUser > 0 {
			return errors.New("email already exists")
		}
		query = query.Set("email = ?", form.Email)
	}
	if form.Name != "" {
		query = query.Set("name = ?", form.Name)
	}

	_, err := query.Exec(ctx)
	return err
}

func (m UserModel) Delete(ctx context.Context, id int64) error {
	tx, err := db.GetDB().Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// delete user groups
	_, err = tx.NewDelete().
		Model((*UserGroup)(nil)).
		Where("user_id = ?", id).
		Exec(ctx)
	if err != nil {
		return err
	}

	// delete user
	_, err = tx.NewDelete().
		Model((*User)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return err
	}

	return tx.Commit()
}
