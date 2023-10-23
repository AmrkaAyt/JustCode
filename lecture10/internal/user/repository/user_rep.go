package repository

import (
	"context"
	"lecture10/internal/user/entity"
)

type UserRepositoryInterface interface {
	GetAll(ctx context.Context) ([]entity.User, error)
	GetById(ctx context.Context, userID int) (*entity.User, error)
	Update(ctx context.Context, userID int, user entity.User) error
	Delete(ctx context.Context, userID int) error
	Register(ctx context.Context, user entity.User, hashedPassword []byte) (int, error)
	Login(ctx context.Context, user entity.User) (*entity.User, error)
}
