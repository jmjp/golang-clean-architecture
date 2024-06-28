package repositories

import (
	"context"
	"onion/internal/domain/entities"
)

type UserRepository interface {
	Save(context context.Context, user *entities.User) (*entities.User, error)
	FindById(context context.Context, email string) (*entities.User, error)
	FindMany(context context.Context, ids []string) ([]*entities.User, error)
	FindWithValidOTP(context context.Context, email, code string) (*entities.User, error)
}
