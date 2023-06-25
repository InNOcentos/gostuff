package dto

type CreateUserDto struct {
	Name 	string `json:"name" validate:"required"`
	Surname string `json:"surname" validate:"required"`
}