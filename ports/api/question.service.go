package api

import (
	"github.com/hoanganh-ng/myhoot/domains/question"
)

type QuestionService struct {
	questionRepo question.Repository
}

func (handler QuestionService) Create(
	dto CreateQuestionRequest,
) (string, error) {
	question := question.New(
		dto.Question,
		dto.WrongAnswers,
		dto.RightAnswer,
	)
	err := handler.questionRepo.Create(question)
	return question.ID(), err
}
