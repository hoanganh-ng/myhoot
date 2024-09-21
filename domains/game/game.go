package game

import (
	"github.com/google/uuid"
)

type Game struct {
	id           uuid.UUID
	listPlayers  []*Player
	listQuestion []*Question
	level        int
}

func (g *Game) ReceiveAnswer() {

}
