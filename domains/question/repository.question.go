package question

type Repository interface {
	Create(*Question) error
	Random(int) ([]Question, error)
}
