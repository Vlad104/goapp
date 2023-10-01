package services

import (
	"app/src/common"
	"app/src/entities"
	"app/src/repositories"
	"log"
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

func (qs *QuestionsService) CurrentCount(availableDto *entities.AvailableQuestionsDto) (int, error) {
	questions, err := qs.questionsRepository.FindAll()
	var currentCount int

	if err != nil {
		return 0, err
	}
	// func countQuestions return count questions(int)
	currentCount = common.MaxQuestionCount - qs.countAvailableQuestions(questions, &availableDto.UserId)

	if common.MaxQuestionCount-currentCount < 0 {
		return 0, nil
	}
	return common.MaxQuestionCount - currentCount, nil
}

func (qs *QuestionsService) countAvailableQuestions(questions []entities.Question, userId *pgtype.UUID) int {

	userQuestions := qs.filterUserId(questions, userId)
	userIntervalQuestions := qs.filterTime(userQuestions)

	return len(userIntervalQuestions)
}

func (qs *QuestionsService) RateLimit(userId *pgtype.UUID) error {
	questions, err := qs.questionsRepository.FindAll()

	if err != nil {
		return err
	}

	count := qs.countAvailableQuestions(questions, userId)

	if count >= common.MaxQuestionCount {
		log.Printf("У пользователя %s превышен порог запросов: %d > %d", common.StringFromUUID(userId), count, common.MaxQuestionCount)
		return common.InternalError
	}

	return nil
}

func (qs *QuestionsService) filterUserId(questions []entities.Question, userId *pgtype.UUID) []entities.Question {
	userQuestions := make([]entities.Question, 0, 0)

	for i := 0; i < len(questions); i++ {

		if common.StringFromUUID(userId) == common.StringFromUUID(&questions[i].UserId) {
			userQuestions = append(userQuestions, questions[i])
		}
	}

	return userQuestions
}

func (qs *QuestionsService) filterTime(userQuestions []entities.Question) []entities.Question {

	userIntervalQuestions := make([]entities.Question, 0, 0)

	for i := 0; i < len(userQuestions); i++ {

		createdAtTime, err := time.Parse(common.SQLTimestampFormatTemplate, userQuestions[i].CreatedAt)

		if err != nil {
			log.Printf("%v", err)
			continue
		}

		if time.Now().Unix()-createdAtTime.Unix() < common.QuestionsRateLimitInterval {

			userIntervalQuestions = append(userIntervalQuestions, userQuestions[i])

		}
	}
	return userIntervalQuestions
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
