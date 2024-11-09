package participant

import (
	"github.com/google/uuid"
	"github.com/hoanganh-ng/myhoot/domains/question"
	customerrors "github.com/hoanganh-ng/myhoot/packages/custom-errors"
)

type Player struct {
	*User
	answerChannel chan<- question.AnswerFromPlayer
}

func (p Player) AnswerQuestion(
	questionID uuid.UUID,
	answer question.Symbol,
) error {
	if p.answerChannel == nil {
		return customerrors.ErrQuestionTimeout
	}
	p.answerChannel <- question.NewAnswerFromPlayer(p.Name(), answer)
	return nil
}

func (p *Player) Listen(
	answerChan chan<- question.AnswerFromPlayer,
	stopAnswerChan <-chan int8,
) {
	p.answerChannel = answerChan
	for range stopAnswerChan {
		p.answerChannel = nil
	}
}
