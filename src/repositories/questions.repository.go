package repositories

import (
	"app/src/common"
	"app/src/database"
	"app/src/entities"
	"context"
	"log"
	"time"
)

type QuestionsRepository struct {
	*database.DataBase
}

func NewQuestionsRepositories(db *database.DataBase) *QuestionsRepository {
	return &QuestionsRepository{db}
}

func (repo *QuestionsRepository) Create(question *entities.CreateQuestionDto) (*entities.Question, error) {
	var id int64
	var createdAt time.Time

	err := repo.DataBase.Conn.QueryRow(
		context.Background(),
		`INSERT INTO "questions" ("userId", "text") VALUES ($1, $2) RETURNING "id", "createdAt"`,
		question.UserID,
		question.Text,
	).Scan(&id, &createdAt)

	if err != nil {
		log.Printf("%v", err)
		return nil, common.InternalError
	}

	// Преобразование в строку с помощью метода Format()
	timeStr := createdAt.Format("2006-01-02 15:04:05")

	result := entities.Question{
		UserId:    question.UserID,
		Text:      question.Text,
		CreatedAt: timeStr,
	}

	return &result, nil
}
