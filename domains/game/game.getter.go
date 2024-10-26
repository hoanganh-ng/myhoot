package game

import (
	"github.com/google/uuid"
)

func (g Game) State() State {
	return g.state
}

func (g Game) ID() uuid.UUID {
	return g.id
}
