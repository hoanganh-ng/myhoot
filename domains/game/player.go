package game

import (
	"github.com/google/uuid"
	"github.com/hoanganh-ng/myhoot/domains/participant"
)

type Player struct {
	*participant.User
	answerChannel chan<- AnswerFromPlayer
}

func (p Player) AnswerQuestion(
	questionID uuid.UUID,
	answer Symbol,
) error {
	if p.answerChannel == nil {
		return ErrQuestionTimeout
	}
	p.answerChannel <- AnswerFromPlayer{p.Name(), answer}
	return nil
}

func (p *Player) Listen(
	answerChan chan<- AnswerFromPlayer,
	stopAnswerChan <-chan int8,
) {
	p.answerChannel = answerChan
	for range stopAnswerChan {
		p.answerChannel = nil
	}
}
