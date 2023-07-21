package services

import (
	"app/src/entities"
	"app/src/repositories"

	"github.com/jackc/pgx/v5/pgtype"
)

// UsersService представляет сервис пользователей.
type UsersService struct {
	repo *repositories.UsersRepository
}

// New создает новый экземпляр UsersService.
func NewUserServices(repo *repositories.UsersRepository) *UsersService {
	return &UsersService{repo}
}

// FindAll возвращает всех пользователей.
func (us *UsersService) FindAll() ([]entities.UserDto, error) {
	users, err := us.repo.FindAll()

	if err != nil {
		return nil, err
	}

	result := make([]entities.UserDto, len(users))

	for i, user := range users {
		result[i] = *toUserDto(&user)
	}

	return result, nil
}

// FindById возвращает пользователя по идентификатору.
func (us *UsersService) FindById(id *pgtype.UUID) (*entities.UserDto, error) {
	user, err := us.repo.FindById(id)

	if err != nil {
		return nil, err
	}

	return toUserDto(user), nil
}

// FindByEmail возвращает пользователя по адресу электронной почты.
func (us *UsersService) FindByEmail(email string) (*entities.UserDto, error) {
	user, err := us.repo.FindByEmail(email)

	if err != nil {
		return nil, err
	}

	return toUserDto(user), nil
}

// Create создает нового пользователя.
func (us *UsersService) Create(createUserDto *entities.CreateUserDto) (*entities.UserDto, error) {
	user, err := us.repo.Create(createUserDto)

	if err != nil {
		return nil, err
	}

	return toUserDto(user), nil
}

// Update обновляет информацию о пользователе.
func (us *UsersService) Update(updateUserDto *entities.UpdateUserDto) (*entities.UserDto, error) {
	user, err := us.repo.Update(updateUserDto)

	if err != nil {
		return nil, err
	}

	return toUserDto(user), nil
}

// Delete удаляет пользователя по идентификатору.
func (us *UsersService) Delete(id *pgtype.UUID) error {
	return us.repo.Delete(id)
}

// toUserDto преобразует сущность User в UserDto.
func toUserDto(user *entities.User) *entities.UserDto {
	return &entities.UserDto{
		ID:    user.ID,
		Email: user.Email,
		Password: user.Password,
	}
}
