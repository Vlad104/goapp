package entities

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID       pgtype.UUID `json:"id"`
	Email    string      `json:"email"`
	Password string      `json:"password"`
}
type UserDto struct {
	ID    pgtype.UUID `json:"id"`
	Email string      `json:"email"`
	Password string
}

type CreateUserDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserDto struct {
	ID       pgtype.UUID `json:"id"`
	Email    string      `json:"email,omitempty"`
	Password string      `json:"password,omitempty"`
}
