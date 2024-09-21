package game

type Answer struct {
	value     string
	isCorrect bool
}

func (a Answer) Value() string {
	return a.value
}

func (a Answer) IsCorrect() bool {
	return a.isCorrect
}
