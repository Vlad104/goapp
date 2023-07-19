package services

import (
	"app/src/entities"
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
		return nil, err
	}

	// Проверяем совпадение паролей
	if user.Password != loginDto.Password {
		return nil, err
	}

	// Создаем AccessToken (логика генерации accessToken будет реализована отдельно)
	accessToken := "user login successful"

	authDto := entities.AuthDto{
		AccessToken: accessToken,
	}

	// Возвращаем AuthDto с AccessToken
	return &authDto, nil
}