package postgres

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `pg:"type:uuid,pk,default:gen_random_uuid()" json:"id"`
	Username string    `pg:"unique:true" json:"username"`
	Password string    `json:"password"`
}
