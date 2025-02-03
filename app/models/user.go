package models

import "time"

type UserRole string

const (
	USER_ADMIN_ROLE   UserRole = "admin"
	USER_DEFAULT_ROLE UserRole = "default"
)

type User struct {
	ID         uint
	FirstName  string
	SecondName string
	Email      string
	Role       UserRole
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}
