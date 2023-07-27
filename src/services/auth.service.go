package services

import (
	"app/src/common"
	"app/src/entities"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewAuthServices(us *UsersService) *AuthService {
	return &AuthService{us}
}

type AuthService struct {
	usersService *UsersService
}

func (authService *AuthService) Login(loginDto *entities.LoginDto) (*entities.AuthDto, error) {
	// Находим пользователя по электронной почте
	user, err := authService.usersService.FindByEmail(loginDto.Email)

	if err != nil {
		return nil, common.ForbiddenError
	}

	// Проверяем совпадение паролей
	if !common.CheckPasswordHash(loginDto.Password, user.Password){
		return nil, common.NotFoundError
	}

	// Создаем AccessToken 
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	"sub": user.ID,
	"iat": time.Now().UnixMilli(),
})

	accessToken, err := token.SignedString(common.SecretKey)
	
	if err != nil {
		log.Printf("%v", err)
		return nil, common.InternalError
	}

	authDto := entities.AuthDto{
		AccessToken: accessToken,
	}
	// Возвращаем AuthDto с AccessToken
	return &authDto, nil
}
