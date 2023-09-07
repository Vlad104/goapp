package entities

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type CreateQuestionDto struct {
	UserID pgtype.UUID `json:"userId"` // идентификатор пользователя задавшего вопрос
	Text   string      `json:"text"`   // текст запроса
}

type Question struct {
	ID        int64       `json:"id"`
	UserId    pgtype.UUID `json:"userId"`
	Text      string      `json:"text"`
	CreatedAt time.Time   `json:"createdAt"`
}
