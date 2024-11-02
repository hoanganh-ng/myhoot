package game

import (
	"github.com/google/uuid"
)

func newGame() *Game {
	newGame := Game{
		id:  uuid.New(),
		pwd: generatePassword(defaultGamePasswordLength),
	}
	newGame.state = CreatedState{&newGame}
	return &newGame
}
