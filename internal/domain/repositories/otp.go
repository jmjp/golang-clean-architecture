package repositories

import (
	"context"
	"onion/internal/domain/entities"
)

type OtpRepository interface {
	Save(context context.Context, otp *entities.OTP) error
	Delete(context context.Context, code string, user string) error
}
