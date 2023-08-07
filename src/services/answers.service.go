package services

import "app/src/entities"

type AnswersService struct {
}

func NewAnswersService() *AnswersService {
	return &AnswersService{}
}

func (as *AnswersService) Create(cq *entities.CreateQuestionDto) (*entities.AnswerDto, error) {
	return &entities.AnswerDto{Text: "Ответ: " + cq.Text}, nil
}
