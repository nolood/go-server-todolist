package postgres

import "github.com/google/uuid"

type User struct {
	ID         uuid.UUID `pg:"type:uuid,pk,default:gen_random_uuid()" json:"id"`
	Username   string    `pg:"unique:true" json:"username"`
	Password   string    `json:"password"`
	TelegramID int64     `json:"telegram_id"`
	VKID       int64     `json:"vk_id"`
	Email      string    `pg:"unique:true" json:"email"`
	Phone      string    `json:"phone"`
	Bills      []*Bill   `pg:"rel:has-many" json:"bills"`
}
