package entities

import (
	"github.com/golang-jwt/jwt/v5"
)
type LoginDto struct{
	Email string `json:"email"`
	Password string `json:"password"`
}

type AuthDto struct{
	AccessToken string `json:"accessToken"`
}

type AuthData struct {
  	UserId string `json:"sub"`
	CreatedAt int64 `json:"iat"` 
  	jwt.RegisteredClaims // техническое поле для парсинга
}
