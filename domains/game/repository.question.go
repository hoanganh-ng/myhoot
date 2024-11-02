package game

type QuestionRepository interface {
	Random(int) ([]Question, error)
}
