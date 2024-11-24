package question

type QuestionService struct {
	questionRepo Repository
}

func (service QuestionService) Create(
	question *Question,
) (string, error) {
	err := service.questionRepo.Create(question)
	return question.ID(), err
}
