package services

import (
	"app/src/common"
	"app/src/entities"
	"app/src/repositories"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type QuestionsService struct {
	answersService      *AnswersService
	questionsRepository *repositories.QuestionsRepository
}

func NewQuestionService(answersService *AnswersService, questionsRepository *repositories.QuestionsRepository) *QuestionsService {
	return &QuestionsService{answersService, questionsRepository}
}

func (qs *QuestionsService) RateLimit(userId *pgtype.UUID) error {
	questions, err := qs.questionsRepository.FindAll()
	if err != nil {
		return err
	}

	userQuestions := make([]entities.Question, 0, 0)
	// Filter userId
	for i := 0; i < len(questions); i++ {

		if common.StringFromUUID(userId) == common.StringFromUUID(&questions[i].UserId){
			userQuestions = append(userQuestions, questions[i])
		}
	}

	userIntervalQuestions := make([]entities.Question, 0, 0)
	// Filter Time
	for i := 0; i < len(userQuestions); i++ {

		createdAtTime, err := time.Parse(common.SQLTimestampFormatTemplate, userQuestions[i].CreatedAt)

		if err != nil {
			return common.InternalError
		}

		if time.Now().Unix()-createdAtTime.Unix() < common.QuestionsRateLimitInterval {

			userIntervalQuestions = append(userIntervalQuestions, questions[i])

		}
	}

	if len(userIntervalQuestions) > common.MaxQuestionCount {
		return common.InternalError
	}

	return nil

}

func (qs *QuestionsService) Create(cq *entities.CreateQuestionDto) (*entities.AnswerDto, error) {
	err := qs.RateLimit(&cq.UserID)
	if err != nil {
		return nil, err
	}
	_, err = qs.questionsRepository.Create(cq)
	if err != nil {
		return nil, err
	}
	answer, err := qs.answersService.Create(cq)
	return answer, err
}
func (qs *QuestionsService) Count() (int, error) {

	question, err := qs.questionsRepository.Count()

	if err != nil {
		return 0, err
	}

	return question, nil

}
