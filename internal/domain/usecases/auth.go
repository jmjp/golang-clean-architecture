package usecases

import (
	"context"
	"onion/internal/domain/entities"
)

type LoginUsecase interface {
	Execute(context context.Context, email string) (*string, error)
}

type VerifyOTPUseCase interface {
	Execute(context context.Context, email, code string) (*entities.User, error)
}

type CreateSessionUseCase interface {
	Execute(context context.Context, identifier string, user string, agent string, ip string) (*entities.Session, error)
}

type GenerateTokenUseCase interface {
	Execute(context context.Context, user *entities.User) (*string, error)
}
