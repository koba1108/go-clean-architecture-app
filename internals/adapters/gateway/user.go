package gateway

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/koba1108/go-clean-architecture-app/internals/domain/model"
	"github.com/koba1108/go-clean-architecture-app/internals/domain/repository"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) GetDB() *sql.DB {
	return u.db
}

func (u *UserRepository) GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	conn := u.GetDB()
	// DB操作は gorm とかに任せたい
	row := conn.QueryRowContext(ctx, "SELECT * FROM `user` WHERE id=?", userID)
	user := model.User{}
	err := row.Scan(&user.ID, &user.FirstName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, errors.New("failed to scan row")
	}
	return &user, nil
}
