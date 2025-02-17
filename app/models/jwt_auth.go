package models

import "time"

type JwtAuth struct {
	ID        uint64
	UserID    uint64
	Token     string
	CreatedAt time.Time
	ExpiresAt time.Time
}
