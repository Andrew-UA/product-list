package models

import "time"

type UserRole string

const (
	USER_ADMIN_ROLE   UserRole = "admin"
	USER_DEFAULT_ROLE UserRole = "default"
)

type User struct {
	ID         uint64
	FirstName  string
	SecondName string
	Email      string
	Nickname   *string
	Role       UserRole
	Password   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}
