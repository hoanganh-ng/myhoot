package game

import (
	"github.com/hoanganhnguyen17/myhoot/domains/user"
)

type Player struct {
	*user.User
	answerChannel chan<- Symbol
}

func (p Player) AnswerQuestion(answer Symbol) {
	p.answerChannel <- answer
}

func (p *Player) Listen(channel chan<- Symbol) {
	p.answerChannel = channel
}
