package services

import (
	"app/src/entities"
	"app/src/repositories"
)

type QuestionsService struct {
	answersService *AnswersService
	questionsRepository	*repositories.QuestionsRepository
}

func NewQuestionService(answersService *AnswersService, questionsRepository	*repositories.QuestionsRepository) *QuestionsService {
	return &QuestionsService{answersService, questionsRepository}
}

func (qs *QuestionsService) Create(cq *entities.CreateQuestionDto) (*entities.AnswerDto, error) {
	_, err := qs.questionsRepository.Create(cq)
	if err != nil {
		return nil, err
	}
	answer, err := qs.answersService.Create(cq)
	return answer, err
}
func(qs *QuestionsService) Count() (int, error) {
	
	question, err := qs.questionsRepository.Count()

	if err != nil {
		return 0, err
	}

	return question, nil

}
