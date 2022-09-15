package gateway

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/koba1108/go-clean-architecture-app/internals/domain/model"
	"github.com/koba1108/go-clean-architecture-app/internals/domain/repository"
	"time"
)

type User struct {
	ID          int       `gorm:"primaryKey;autoIncrement:not null"`
	DisplayName string    `gorm:"not null"`
	FirstName   string    `gorm:"not null"`
	LastName    string    `gorm:"not null"`
	Birthday    time.Time `gorm:"not null;type:date"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
	DeletedAt   *time.Time
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetDB() *gorm.DB {
	return r.db
}

func (r *UserRepository) GetAll() ([]*model.User, error) {
	conn := r.GetDB()
	var users []*User
	conn.Find(&users)
	if conn.Error != nil {
		return nil, conn.Error
	}
	return r.toUserModels(users), nil
}

func (r *UserRepository) GetByID(userID int) (*model.User, error) {
	conn := r.GetDB()
	var user model.User
	conn.First(&user, userID)
	if conn.Error != nil {
		if errors.Is(conn.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, conn.Error
	}
	return &user, nil
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	conn := r.GetDB()
	user := r.toUserEntity(u)
	conn.Create(user)
	if conn.Error != nil {
		return nil, conn.Error
	}
	return r.toUserModel(user), nil
}

func (r *UserRepository) Update(u *model.User) (*model.User, error) {
	conn := r.GetDB()
	user := r.toUserEntity(u)
	conn.Save(user)
	if conn.Error != nil {
		return nil, conn.Error
	}
	return r.toUserModel(user), nil
}

func (r *UserRepository) DeleteByID(userID int) error {
	conn := r.GetDB()
	conn.Delete(&User{}, userID)
	if conn.Error != nil {
		return conn.Error
	}
	return nil
}

func (r *UserRepository) toUserModel(u *User) *model.User {
	return &model.User{
		ID:              u.ID,
		DisplayName:     u.DisplayName,
		FirstName:       u.FirstName,
		LastName:        u.LastName,
		FullName:        fmt.Sprintf("%s %s", u.LastName, u.FirstName),
		Birthday:        u.Birthday,
		Age:             int(time.Now().Sub(u.Birthday).Hours() / 24 / 365),
		RegisteredAt:    u.CreatedAt,
		LatestUpdatedAt: u.UpdatedAt,
	}
}

func (r *UserRepository) toUserModels(us []*User) []*model.User {
	var users []*model.User
	for _, u := range us {
		users = append(users, r.toUserModel(u))
	}
	return users
}

func (r *UserRepository) toUserEntity(u *model.User) *User {
	return &User{
		ID:          u.ID,
		DisplayName: u.DisplayName,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Birthday:    u.Birthday,
		CreatedAt:   u.RegisteredAt,
		UpdatedAt:   u.LatestUpdatedAt,
	}
}

func (r *UserRepository) toUserEntities(us []*model.User) []*User {
	var users []*User
	for _, u := range us {
		users = append(users, r.toUserEntity(u))
	}
	return users
}
