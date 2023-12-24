package models

type CreateUserDto struct {
	tableName struct{} `pg:"users"`
	Username  string   `json:"username"`
	Password  string   `json:"password"`
}
