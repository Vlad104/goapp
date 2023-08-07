package services

import (
	"app/src/entities"
)

// QuestionService.Create // вызывает AnswerService для генерации ответа
// (в будущем будет добавлена логика по сохранению вопросов в базу данных)

type QuestionService struct {
	answersService *AnswersService
}

func NewQuestionService(answersService *AnswersService) *QuestionService {
	return &QuestionService{answersService}
}

func (qs *QuestionService) Create(cq *entities.CreateQuestionDto) (*entities.AnswerDto, error) {
	answer, err := qs.answersService.Create(cq)
	return answer, err
}
