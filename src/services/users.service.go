package services

import (
	"app/src/entities"
	"app/src/repositories"

	"github.com/jackc/pgx/v5/pgtype"
)

type UsersService struct {
	repo *repositories.UsersRepository
}

func New(repo *repositories.UsersRepository) *UsersService {
	return &UsersService{repo}
}

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

func (service *UsersService) FindById(id *pgtype.UUID) (*entities.UserDto, error) {
	user, err := service.repo.FindById(id)

	if err != nil {
		return nil, err
	}

	return toUserDto(user), nil
}

func (service *UsersService) Create(createUserDto *entities.CreateUserDto) (*entities.UserDto, error) {
	user, err := service.repo.Create(createUserDto)

	if err != nil {
		return nil, err
	}

	return toUserDto(user), nil
}

func (service *UsersService) Update(updateUserDto *entities.UpdateUserDto) (*entities.UserDto, error) {
	user, err := service.repo.Update(updateUserDto)

	if err != nil {
		return nil, err
	}

	return toUserDto(user), nil
}

func (service *UsersService) Delete(id *pgtype.UUID) error {
	return service.repo.Delete(id)
}

func toUserDto(user *entities.User) *entities.UserDto {
	return &entities.UserDto{
		ID:    user.ID,
		Email: user.Email,
	}
}
