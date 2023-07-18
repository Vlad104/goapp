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
func New(repo *repositories.UsersRepository) *UsersService {
	return &UsersService{repo}
}

// FindAll возвращает всех пользователей.
func (service *UsersService) FindAll() ([]entities.UserDto, error) {
	users, err := service.repo.FindAll()

	if err != nil {
		return nil, err
	}

	result := make([]entities.UserDto, len(users))

	for i, user := range users {
		result[i] = entities.UserDto{
			ID:    user.ID,
			Email: user.Email,
		}
	}

	return result, nil
}

// FindById возвращает пользователя по идентификатору.
func (service *UsersService) FindById(id *pgtype.UUID) (*entities.UserDto, error) {
	user, err := service.repo.FindById(id)

	if err != nil {
		return nil, err
	}

	return toUserDto(user), nil
}

// FindByEmail возвращает пользователя по адресу электронной почты.
func (service *UsersService) FindByEmail(email string) (*entities.UserDto, error) {
	user, err := service.repo.FindByEmail(email)

	if err != nil {
		return nil, err
	}

	return toUserDto(user), nil
}

// Create создает нового пользователя.
func (service *UsersService) Create(createUserDto *entities.CreateUserDto) (*entities.UserDto, error) {
	user, err := service.repo.Create(createUserDto)

	if err != nil {
		return nil, err
	}

	return toUserDto(user), nil
}

// Update обновляет информацию о пользователе.
func (service *UsersService) Update(updateUserDto *entities.UpdateUserDto) (*entities.UserDto, error) {
	user, err := service.repo.Update(updateUserDto)

	if err != nil {
		return nil, err
	}

	return toUserDto(user), nil
}

// Delete удаляет пользователя по идентификатору.
func (service *UsersService) Delete(id *pgtype.UUID) error {
	return service.repo.Delete(id)
}

// toUserDto преобразует сущность User в UserDto.
func toUserDto(user *entities.User) *entities.UserDto {
	return &entities.UserDto{
		ID:    user.ID,
		Email: user.Email,
	}
}
