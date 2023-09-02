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

func(repo *QuestionsRepository) Count() (int, error) {
  var count int
    err := repo.DataBase.Conn.QueryRow(
        context.Background(),
        `SELECT COUNT(*) FROM "questions"`,
    ).Scan(&count)

    if err != nil {
        log.Printf("%v", err)
        return 0, common.InternalError
    }
    return count, nil
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
	timeStr := createdAt.Format(common.SQLTimestampFormatTemplate)

	result := entities.Question{
		UserId:    question.UserID,
		Text:      question.Text,
		CreatedAt: timeStr,
	}

	return &result, nil
}
