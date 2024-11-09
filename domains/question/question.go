package question

import (
	"math/rand"

	"github.com/google/uuid"
)

type Question struct {
	id         uuid.UUID
	value      string
	listAnswer []Answer
}

func New(
	question string,
	wrongAnswers []string,
	rightAnswer string,
) *Question {
	listAnswer := make([]Answer, 0)
	for i := 0; i < len(wrongAnswers); i++ {
		asw := NewAnswer(wrongAnswers[i], false)
		listAnswer = append(listAnswer, asw)
	}
	listAnswer = append(listAnswer, NewAnswer(rightAnswer, true))
	return &Question{
		id:         uuid.New(),
		value:      question,
		listAnswer: listAnswer,
	}
}

func (q Question) ID() string {
	return q.id.String()
}

func (q Question) Value() string {
	return q.value
}

func (q Question) rightAnser() Answer {
	var rightAnswer Answer
	for i := 0; i < len(q.listAnswer); i++ {
		if q.listAnswer[i].IsCorrect() {
			rightAnswer = q.listAnswer[i]
		}
	}
	return rightAnswer
}

func (q Question) CompareAnswer(answer Symbol) bool {
	rightAnswer := q.rightAnser()
	return *rightAnswer.symbol == answer
}

func (q Question) ListAnswer(shuffle bool) []Answer {
	if !shuffle {
		return q.listAnswer
	}
	mcheck := make(map[int]bool)
	shuffledList := make([]Answer, 0)
	max := len(q.listAnswer)
	for i := 0; i < max; i++ {
	dubCheck:
		for {
			r := rand.Intn(max)
			if _, ok := mcheck[r]; !ok {
				mcheck[r] = true
				shuffledList = append(shuffledList, q.listAnswer[r])
				break dubCheck
			}
		}
	}
	return shuffledList
}
