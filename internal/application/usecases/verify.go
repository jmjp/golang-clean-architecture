package usecases

import (
	"context"
	"onion/internal/domain/entities"
	"onion/internal/domain/repositories"
)

type VerifyOTPUseCase struct {
	user repositories.UserRepository
	otp  repositories.OtpRepository
}

// NewVerifyOTPUseCase creates a new instance of the VerifyOTPUseCase struct.
//
// It takes a userRepository of type repositories.UserRepository, an otpRepository of type repositories.OtpRepository, and a sessionRepository of type repositories.SessionRepository as parameters.
// It returns a pointer to the VerifyOTPUseCase struct.
func NewVerifyOTPUseCase(user repositories.UserRepository, otp repositories.OtpRepository, session repositories.SessionRepository) *VerifyOTPUseCase {
	return &VerifyOTPUseCase{
		user: user,
		otp:  otp,
	}
}

// Execute is a method of the VerifyOTPUseCase struct that executes the OTP verification process.
//
// It takes a context.Context, email, and code as parameters. The context is used for control during execution, email represents the user's email, and code is the OTP code to verify.
// It returns a pointer to an entities.User and an error.
func (u *VerifyOTPUseCase) Execute(ctx context.Context, email, code string) (*entities.User, error) {
	user, err := u.user.FindWithValidOTP(ctx, email, code)
	if err != nil {
		return nil, err
	}
	go func() {
		err := u.otp.Delete(context.Background(), code, email)
		if err != nil {
			panic(err)
		}
	}()
	return user, nil
}
