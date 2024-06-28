package entities

import "time"

type Session struct {
	ID         int       `json:"id"`
	Identifier string    `json:"identifier"`
	ValidUntil time.Time `json:"valid_until"`
	Agent      string    `json:"agent"`
	User       string    `json:"user"`
	IP         string    `json:"ip"`
}

func NewSession(identifier string, user string, agent string, ip string) *Session {
	return &Session{
		Identifier: identifier,
		ValidUntil: time.Now().Add(15 * time.Minute),
		Agent:      agent,
		User:       user,
		IP:         ip,
	}
}
