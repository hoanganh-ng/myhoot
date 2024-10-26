package game

type Symbol uint8

const (
	Triangle = iota
	Square
	Circle
	Rectangle
)

type AnswerFromPlayer struct {
	playerName string
	symbol     Symbol
}

type Answer struct {
	value     string
	isCorrect bool
	Symbol    Symbol
}

func (a Answer) Value() string {
	return a.value
}

func (a Answer) IsCorrect() bool {
	return a.isCorrect
}
