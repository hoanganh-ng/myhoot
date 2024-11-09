package question

type QuestionRepository interface {
	Random(int) ([]Question, error)
}
