package dto

type UserDTO struct {
	FirstName  *string
	SecondName *string
	Email      *string
	Role       *string
	Nickname   Nullable[string]
	Password   Nullable[string]
}
