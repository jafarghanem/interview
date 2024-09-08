package storage

import (
	"context"
	"errors"
	"strings"
	"users/internal/user"
	"users/pkg/adapters/storage/entities"
	"users/pkg/adapters/storage/mappers"
	"users/pkg/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)
type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) user.Repo {
	return &userRepo{
		db: db,
	}
}
func (r *userRepo) Create(ctx context.Context, user *user.User) (*user.User, error) {
	newUser := mappers.UserDomainToEntity(*user)
	err := r.db.WithContext(ctx).Create(&newUser).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, utils.DbErrDuplicateKey
		}
		return nil, err
	}
	createdUser := mappers.UserEntityToDomain(newUser)
	return &createdUser, nil
}

func (r *userRepo) GetByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
	var u entities.User

	err := r.db.WithContext(ctx).Model(&entities.User{}).Where("id = ?", id).First(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	fetchedUser := mappers.UserEntityToDomain(u)
	return &fetchedUser, nil
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	var user entities.User
	err := r.db.WithContext(ctx).Model(&entities.User{}).Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	fetchedUsers := mappers.UserEntityToDomain(user)
	return &fetchedUsers, nil
}
