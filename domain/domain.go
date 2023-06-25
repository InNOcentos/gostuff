package domain

type User struct {
	ID 			string `json:"id"`
	Name 		string `json:"name" validate:"required"`
	Surname string `json:"surname" validate:"required"`
}