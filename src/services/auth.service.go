package services

import (
	"app/src/entities"
	"app/src/common"

	"golang.org/x/crypto/bcrypt"
)

func NewAuthServices(us *UsersService) *AuthService {
	return &AuthService{us}
}

type AuthService struct {
	usersService *UsersService
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func (authService *AuthService) Login(loginDto *entities.LoginDto) (*entities.AuthDto, error) {
	// Находим пользователя по электронной почте
	user, err := authService.usersService.FindByEmail(loginDto.Email)

	if err != nil {
		return nil, common.ForbiddenError
	}

	// Проверяем совпадение паролей
	if !CheckPasswordHash(loginDto.Password, user.Password){
		return nil, common.NotFoundError
	}

	// Создаем AccessToken (логика генерации accessToken будет реализована отдельно)
	accessToken := "user login successful"

	authDto := entities.AuthDto{
		AccessToken: accessToken,
	}

	// Возвращаем AuthDto с AccessToken
	return &authDto, nil
}
