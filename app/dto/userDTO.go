package dto

type UserDTO struct {
	FirstName  *string
	SecondName *string
	Email      *string
	Nickname   Nullable[string]
}
