package game

import "github.com/hoanganhnguyen17/myhoot/domains/participant"

type Player struct {
	*participant.User
	answerChannel chan<- Symbol
}

func (p Player) AnswerQuestion(answer Symbol) {
	p.answerChannel <- answer
}
func (p Player) Is(other any) bool {
	otherPlayer, ok := other.(Player)
	if !ok {
		return false
	}
	return p.Name() == otherPlayer.Name()
}

func (p *Player) Listen(channel chan<- Symbol) {
	p.answerChannel = channel
}
