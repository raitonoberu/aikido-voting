package relations

import (
	"aikido/db"
	"aikido/models"
	"context"
	"database/sql"
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

	for _, m := range []any{
		(*models.User)(nil),
		(*models.Pool)(nil),
		(*models.Option)(nil),
		(*models.Group)(nil),
		(*models.UserGroup)(nil),
	} {
		_, err := tx.NewCreateTable().
			IfNotExists().
			Model(m).
			Exec(ctx)
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}
