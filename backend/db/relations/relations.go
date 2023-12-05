package relations

import (
	"aikido/db"
	"aikido/models"
	"context"
	"database/sql"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func Create() error {
	db := db.GetDB()
	ctx := context.Background()

	db.RegisterModel((*models.UserGroup)(nil))

	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// insert models
	for _, m := range []any{
		(*models.User)(nil),
		(*models.Pool)(nil),
		(*models.Option)(nil),
		(*models.Group)(nil),
		(*models.UserGroup)(nil),
		(*models.Vote)(nil),
		(*models.Message)(nil),
	} {
		_, err := tx.NewCreateTable().
			IfNotExists().
			Model(m).
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	// insert default groups
	for _, g := range []models.Group{
		{ID: 1, Name: "Преподаватели"},
		{ID: 2, Name: "Все пользователи"},
	} {
		_, err := tx.NewInsert().
			Model(&g).
			On("conflict (id) do nothing").
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	// insert default admin
	adminEmail, adminPassword := os.Getenv("ADMIN_EMAIL"), os.Getenv("ADMIN_PASSWORD")
	hashedPassword, _ := bcrypt.GenerateFromPassword(
		[]byte(adminPassword), bcrypt.DefaultCost,
	)

	u := &models.User{
		ID:       1,
		RoleID:   1,
		Name:     "Администратор",
		Email:    adminEmail,
		Password: string(hashedPassword),
	}
	ug := &models.UserGroup{
		UserID:  1,
		GroupID: 1,
	}

	_, err = tx.NewInsert().
		Model(u).
		On("conflict (id) do nothing").
		Exec(ctx)
	if err != nil {
		return err
	}
	_, err = tx.NewInsert().
		Model(ug).
		On("conflict (user_id, group_id) do nothing").
		Exec(ctx)
	if err != nil {
		return err
	}

	return tx.Commit()
}
