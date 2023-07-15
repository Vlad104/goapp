package repositories

import (
	"app/src/common"
	"app/src/database"
	"app/src/entities"
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgtype"
)

type UsersRepository struct {
	*database.DataBase
}

func New(pg *database.DataBase) *UsersRepository {
	return &UsersRepository{pg}
}

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

func (repo *UsersRepository) Create(user *entities.CreateUserDto) (*entities.User, error) {
	var id pgtype.UUID

	err := repo.DataBase.Conn.QueryRow(
		context.Background(),
		`INSERT INTO "users" (email, password) VALUES ($1, $2) RETURNING "id"`,
		user.Email,
		user.Password,
	).Scan(&id)

	if err != nil {
		log.Printf("%v", err)
		return nil, common.InternalError
	}

	result := entities.User{
		ID:       id,
		Email:    user.Email,
		Password: user.Password,
	}

	return &result, nil
}

func (repo *UsersRepository) Update(user *entities.UpdateUserDto) (*entities.User, error) {
	_, err := repo.DataBase.Conn.Exec(
		context.Background(),
		`UPDATE "users" SET "email" = $2, "password" = $3 WHERE "id" = $1`,
		user.ID,
		user.Email,
		user.Password,
	)

	if err != nil {
		log.Printf("%v", err)
		return nil, common.InternalError
	}

	result := entities.User(*user)

	return &result, nil
}

func (repo *UsersRepository) Delete(id *pgtype.UUID) error {
	_, err := repo.DataBase.Conn.Exec(
		context.Background(),
		`DELETE FROM "users" WHERE "id" = $1`,
		*id,
	)

	return err
}
