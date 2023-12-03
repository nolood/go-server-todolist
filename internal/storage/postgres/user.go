package postgres

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `pg:"type:uuid,pk,default:gen_random_uuid()" json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
}
