package game

import (
	"github.com/google/uuid"
	"github.com/hoanganhnguyen17/myhoot/helpers"
)

type Participant interface {
	ReceiveLeaderboard()
}

type Game struct {
	id              uuid.UUID
	state           State
	manager         Manager
	currentQuestion int
	players         []Player
	questions       []Question
	leaderboard     Leaderboard
}

func New() *Game {
	newGame := Game{
		id: uuid.New(),
	}
	newGame.state = CreatedState{&newGame}
	return &newGame
}

func (game *Game) onStart() error {
	err := game.state.Started()
	if err == nil {
		game.currentQuestion = 1
	}
	return err
}

func (game *Game) startReceiveAnswer() error {
	return nil
}

func (game *Game) PlayerJoin(player Player) error {
	if !game.state.Is(Created) {
		return ErrPlayerCannotJoinTheGame
	}
	var added bool
	game.players, added = helpers.AddIfNotExist(game.players, player)
	if !added {
		return ErrPlayerIsJoined
	}
	return nil
}

func (game *Game) setQuestionList(
	input []Question,
) error {
	if !game.state.Is(Created) {
		return ErrCannotSetListOfQuestion
	}
	game.questions = input
	return nil
}

func (game *Game) onNextQuestion() error {
	if game.currentQuestion == len(game.questions) {
		return ErrCannotGoToTheNextQuestion
	}
	game.currentQuestion += 1
	return nil
}

func (game Game) CurrentQuestion() (Question, error) {
	var q Question
	if !game.state.Is(Started) {
		return q, ErrGameIsNotStartYet
	}
	q = game.questions[game.currentQuestion-1]
	return q, nil
}

func (game *Game) updateState(newState State) {
	game.state = newState
}
