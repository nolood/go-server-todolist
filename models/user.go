package models

type CreateUserDto struct {
	tableName struct{} `pg:"users"`
	Username  string   `json:"username"`
	Password  string   `json:"password"`
}

type CheckUserDto struct {
	tableName struct{} `pg:"users"`
	VKID      int64    `json:"vk_id"`
	Username  string   `json:"username"`
	Password  string   `json:"password"`
}
