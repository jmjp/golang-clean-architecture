package entities

import (
	"onion/pkg/random"
	"time"
)

type OTP struct {
	ID         int       `json:"id"`
	Code       string    `json:"code"`
	User       string    `json:"user"`
	ValidUntil time.Time `json:"valid_until"`
}

func NewOTP(user string) *OTP {
	return &OTP{
		Code:       random.Int(4),
		User:       user,
		ValidUntil: time.Now().Add(15 * time.Minute),
	}
}
