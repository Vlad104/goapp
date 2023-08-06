package entities

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateQuestionDto struct {
	UserID   pgtype.UUID `json:"userId"` // идентификатор пользователя задавшего вопрос
	Text     string      `json:"text"`   // текст запроса
}

