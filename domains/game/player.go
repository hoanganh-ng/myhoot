package game

import (
	"github.com/google/uuid"
	"github.com/hoanganhnguyen17/myhoot/domains/participant"
)

type Player struct {
	*participant.User
	answerChannel chan<- AnswerFromPlayer
}

func (p Player) AnswerQuestion(
	questionID uuid.UUID,
	answer Symbol,
) {
	p.answerChannel <- AnswerFromPlayer{p.Name(), answer}
}
func (p Player) Is(other any) bool {
	otherPlayer, ok := other.(Player)
	if !ok {
		return false
	}
	return p.Name() == otherPlayer.Name()
}

func (p *Player) Listen(
	answerChan chan<- AnswerFromPlayer,
	stopAnswerChan <-chan int8,
) {
	p.answerChannel = answerChan
	for {
		select {
		case <-stopAnswerChan:
			p.answerChannel = nil
			return
		}
	}
}
