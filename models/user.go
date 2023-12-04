package models

type CreateUserDto struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
