package usecases

import (
	"context"
	"onion/internal/domain/entities"
	"onion/internal/domain/repositories"
)

type CreateSessionUsecase struct {
	sessionRepository repositories.SessionRepository
}

// NewCreateSessionUsecase creates a new instance of the CreateSessionUsecase struct.
//
// It takes a sessionRepository of type repositories.SessionRepository as a parameter.
// It returns a pointer to the CreateSessionUsecase struct.
func NewCreateSessionUsecase(sessionRepository repositories.SessionRepository) *CreateSessionUsecase {
	return &CreateSessionUsecase{
		sessionRepository: sessionRepository,
	}
}

// Execute executes the CreateSessionUsecase function.
//
// It takes context.Context, identifier string, user string, agent string, and ip string as parameters.
// It returns a pointer to entities.Session and an error.
func (uc *CreateSessionUsecase) Execute(ctx context.Context, identifier string, user string, agent string, ip string) (*entities.Session, error) {
	session := entities.NewSession(identifier, user, agent, ip)
	err := uc.sessionRepository.Save(ctx, session)
	if err != nil {
		return nil, err
	}
	return session, nil
}
