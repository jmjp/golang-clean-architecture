package usecases

import (
	"context"
	"onion/internal/domain/entities"
	"onion/internal/domain/repositories"
)

type LoginUsecase struct {
	user repositories.UserRepository
	otp  repositories.OtpRepository
}

// NewLoginUsecase creates a new instance of the LoginUsecase struct.
//
// It takes a userRepository of type repositories.UserRepository and an otpRepository of type repositories.OtpRepository as parameters.
// It returns a pointer to the LoginUsecase struct.
func NewLoginUsecase(user repositories.UserRepository, otp repositories.OtpRepository) *LoginUsecase {
	return &LoginUsecase{
		user: user,
		otp:  otp,
	}
}

// Execute is a method of the LoginUsecase struct that performs a login operation.
//
// It takes a context.Context and an email string as parameters. The context is used to control the execution of the function, and the email parameter is the user's email address.
// The function returns a pointer to a string and an error. The string is the OTP code generated for the user, and the error is any error that occurred during the execution of the function.
func (u *LoginUsecase) Execute(ctx context.Context, email string) (*string, error) {
	userEntity, err := entities.NewUserEntity(email)
	if err != nil {
		return nil, err
	}
	user, err := u.user.Save(ctx, userEntity)
	if err != nil {
		return nil, err
	}
	otpEntity := entities.NewOTP(user.ID)
	go func() {
		err := u.otp.Save(context.Background(), otpEntity)
		if err != nil {
			panic(err)
		}
	}()
	return &otpEntity.Code, nil
}
