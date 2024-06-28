package entities

import (
	"errors"
	"net/mail"
	"onion/pkg/random"
	"regexp"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Avatar    *string   `json:"avatar"`
	Username  string    `json:"username"`
	Blocked   bool      `json:"blocked"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUserEntity(email string) (*User, error) {
	user := &User{
		Email:     email,
		Username:  "hyperzoop_" + random.String(5),
		Blocked:   false,
		Avatar:    nil,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := user.isValid(); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) isValid() error {
	username := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	if !username.MatchString(u.Username) && len(u.Username) < 6 {
		return errors.New("username must be at least 6 characters long and contain only alphanumeric characters and underscores")
	}
	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return errors.New("invalid email")
	}
	if u.Avatar != nil && !regexp.MustCompile(`^https?://`).MatchString(*u.Avatar) {
		return err
	}
	return nil
}
