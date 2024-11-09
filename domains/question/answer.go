package question

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

func NewAnswerFromPlayer(
	playerName string,
	answerSymbol Symbol,
) AnswerFromPlayer {
	return AnswerFromPlayer{playerName, answerSymbol}
}

func (ap AnswerFromPlayer) PlayerName() string {
	return ap.playerName
}
func (ap AnswerFromPlayer) Symbol() Symbol {
	return ap.symbol
}

type Answer struct {
	value     string
	isCorrect bool
	symbol    *Symbol
}

func NewAnswer(answer string, isCorrect bool) Answer {
	return Answer{value: answer, isCorrect: isCorrect}
}

func (a *Answer) setSymbol(input *Symbol) {
	a.symbol = input
}

func (a Answer) Value() string {
	return a.value
}

func (a Answer) IsCorrect() bool {
	return a.isCorrect
}
