package repositories

import (
	"app/src/common"
	"app/src/database"
	"app/src/entities"
	"context"
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/jackc/pgx/v5/pgtype"
)

// UsersRepository представляет репозиторий пользователей.
type UsersRepository struct {
	*database.DataBase
}

// New создает новый экземпляр UsersRepository.
func NewUserRepositories(pg *database.DataBase) *UsersRepository {
	return &UsersRepository{pg}
}

// FindAll возвращает всех пользователей.
func (repo *UsersRepository) FindAll() ([]entities.User, error) {
	rows, err := repo.DataBase.Conn.Query(
		context.Background(),
		`SELECT * FROM "users"`,
	)
	defer rows.Close()

	if err != nil {
		log.Printf("%v", err)
		return nil, common.InternalError
	}

	users := make([]entities.User, 0, 0)

	for rows.Next() {
		user := entities.User{}

		err = rows.Scan(
			&user.ID,
			&user.Email,
			&user.Password,
		)

		if err != nil {
			log.Printf("%v", err)
			return nil, common.InternalError
		}

		users = append(users, user)
	}

	return users, nil
}

// FindById возвращает пользователя по идентификатору.
func (repo *UsersRepository) FindById(id *pgtype.UUID) (*entities.User, error) {
	user := entities.User{}

	err := repo.DataBase.Conn.QueryRow(
		context.Background(),
		`SELECT * FROM "users" WHERE "id" = $1`,
		*id,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		log.Printf("%v", err)
		return nil, common.NotFoundError
	}

	return &user, nil
}

// FindByEmail возвращает пользователя по адресу электронной почты.
func (repo *UsersRepository) FindByEmail(email string) (*entities.User, error) {
	user := entities.User{}

	err := repo.DataBase.Conn.QueryRow(
		context.Background(),
		`SELECT * FROM "users" WHERE "email" = $1`,
		email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		log.Printf("%v", err)
		return nil, common.NotFoundError
	}

	return &user, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Create создает нового пользователя.
func (repo *UsersRepository) Create(user *entities.CreateUserDto) (*entities.User, error) {
	var id pgtype.UUID
	hash, err := HashPassword(user.Password)
	
	if err != nil{
		log.Printf("%v",err)
		return nil, common.InternalError
	}

	err = repo.DataBase.Conn.QueryRow(
		context.Background(),
		`INSERT INTO "users" (email, password) VALUES ($1, $2) RETURNING "id"`,
		user.Email,
		hash,
	).Scan(&id)

	if err != nil {
		log.Printf("%v", err)
		return nil, common.InternalError
	}

	result := entities.User{
		ID:       id,
		Email:    user.Email,
		Password: hash,
	}

	return &result, nil
}

// Update обновляет информацию о пользователе.
func (repo *UsersRepository) Update(user *entities.UpdateUserDto) (*entities.User, error) {
	hash, err := HashPassword(user.Password)

	_, err = repo.DataBase.Conn.Exec(
		context.Background(),
		`UPDATE "users" SET "email" = $2, "password" = $3 WHERE "id" = $1`,
		user.ID,
		user.Email,
		hash,
	)

	if err != nil {
		log.Printf("%v", err)
		return nil, common.InternalError
	}

	result := entities.User(*user)

	return &result, nil
}

// Delete удаляет пользователя по идентификатору.
func (repo *UsersRepository) Delete(id *pgtype.UUID) error {
	_, err := repo.DataBase.Conn.Exec(
		context.Background(),
		`DELETE FROM "users" WHERE "id" = $1`,
		*id,
	)

	return err
}
