package game

import (
	"github.com/google/uuid"
	"github.com/hoanganhnguyen17/myhoot/domains/user"
)

type Participant interface {
	ReceiveLeaderboard()
}

type Game struct {
	id           uuid.UUID
	admin        *user.User
	listPlayers  []*Player
	state        State
	listQuestion []*Question
	level        int
}

func (g *Game) updateState(newState State) {
	g.state = newState
}
func (g *Game) ReceiveAnswer() {

}
