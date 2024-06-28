package repositories

import (
	"context"
	"onion/internal/domain/entities"
)

type SessionRepository interface {
	Save(context context.Context, session *entities.Session) error
	FindOne(context context.Context, identifier string) (*entities.Session, error)
	FindMany(context context.Context, userId string) ([]*entities.Session, error)
	Delete(context context.Context, identifier string) error
}
