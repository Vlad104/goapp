package controllers

import (
	"app/src/common"
	"app/src/entities"
	"app/src/services"
	"encoding/json"
	"net/http"
)

type QuestionsController struct {
	service *services.QuestionsService
}

func NewQuestionsController(service *services.QuestionsService) *QuestionsController {
	return &QuestionsController{service}
}

func (qs *QuestionsController) Count(w http.ResponseWriter, r *http.Request) {
	questions, err := qs.service.Count()

	if err != nil {
		common.HandleHttpError(w, err)
		return
	}

	response, err := json.Marshal(questions)

	if err != nil {
		common.HandleHttpError(w, err)
		return
	}

	w.Write(response)
}

func (qc *QuestionsController) Create(w http.ResponseWriter, r *http.Request) {
	var createQuestionDto entities.CreateQuestionDto

	err := json.NewDecoder(r.Body).Decode(&createQuestionDto)

	if err != nil {
		common.HandleHttpError(w, err)
		return
	}

	question, err := qc.service.Create(&createQuestionDto)

	if err != nil {
		common.HandleHttpError(w, err)
		return
	}

	response, err := json.Marshal(question)

	if err != nil {
		common.HandleHttpError(w, err)
		return
	}

	w.Write(response)
}

//QuestionsController.Create  парсит параметры запроса "/questions"
