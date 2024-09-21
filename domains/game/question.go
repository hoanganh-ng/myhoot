package game

import (
	"math/rand"

	"github.com/google/uuid"
)

type Question struct {
	id         uuid.UUID
	value      string
	listAnswer []Answer
}

func (q Question) ID() string {
	return q.id.String()
}

func (q Question) Value() string {
	return q.value
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
