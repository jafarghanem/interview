package user

import (
	"context"
	"errors"
	"users/pkg/utils"

	"github.com/google/uuid"
)

type Ops struct {
	repo Repo
}

func NewOps(repo Repo) *Ops {
	return &Ops{repo}
}

func (o *Ops) Create(ctx context.Context, user *User) (*User, error) {

	err := validateUserRegistration(user)
	if err != nil {
		return nil, err
	}
	hashedPass, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.SetPassword(hashedPass)

	// lowercase email
	user.Email = LowerCaseEmail(user.Email)
	createdUser, err := o.repo.Create(ctx, user)
	if err != nil {
		if errors.Is(err, utils.DbErrDuplicateKey) {
			return nil, ErrEmailAlreadyExists
		}
		return nil, err
	}
	return createdUser, nil
}

func (o *Ops) GetUserByID(ctx context.Context, id uuid.UUID) (*User, error) {
	return o.repo.GetByID(ctx, id)
}

func (o *Ops) GetUserByEmailAndPassword(ctx context.Context, email, password string) (*User, error) {
	email = LowerCaseEmail(email)
	user, err := o.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrUserNotFound
	}

	if err := utils.CheckPasswordHash(password, user.Password); err != nil {
		return nil, ErrInvalidAuthentication
	}

	return user, nil
}

func (o *Ops) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	email = LowerCaseEmail(email)
	user, err := o.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func validateUserRegistration(user *User) error {
	err := ValidateEmail(user.Email)
	if err != nil {
		return err
	}

	if err := ValidatePasswordWithFeedback(user.Password); err != nil {
		return err
	}
	return nil
}