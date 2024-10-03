package game

import (
	"github.com/google/uuid"
)

type Participant interface {
	ReceiveLeaderboard()
}

type Game struct {
	id              uuid.UUID
	state           State
	manager         Manager
	players         []Player
	questions       []Question
	currentQuestion int
	leaderboard     Leaderboard
}

func (g *Game) Started(newState State) {
	g.state.Started()
}

func (g *Game) SendQuestion() {

}

func (g Game) State() State {
	return g.state
}

func (g Game) ID() uuid.UUID {
	return g.id
}

func (g Game) CurrentQuestion() Question {
	return g.questions[g.currentQuestion]
}

func (g *Game) updateState(newState State) {
	g.state = newState
}
