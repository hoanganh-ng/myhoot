package game

import (
	"time"

	"github.com/google/uuid"
	"github.com/hoanganh-ng/myhoot/domains/participant"
	"github.com/hoanganh-ng/myhoot/domains/question"
	customerrors "github.com/hoanganh-ng/myhoot/packages/custom-errors"
)

const (
	rightAnswerMaxPoint       uint32 = 5 * 1000 //max is reached when player anwser in 0 miliseconds
	defaultGamePasswordLength int    = 5
)

type Participant interface {
	ReceiveLeaderboard()
}

type Game struct {
	id              uuid.UUID
	state           State
	manager         *Manager
	currentQuestion int
	players         map[string]*participant.Player
	questions       []question.Question
	leaderboard     Leaderboard
	pwd             string
}

func (game Game) Authenticate(pwd string) bool {
	return game.pwd == pwd
}

func (game Game) CurrentQuestion() (question.Question, error) {
	var q question.Question
	if !game.state.Is(Started) {
		return q, customerrors.ErrGameIsNotStartYet
	}
	q = game.questions[game.currentQuestion-1]
	return q, nil
}

func (game *Game) OnPlayerRequestJoin(player *participant.Player) error {
	if !game.state.Is(Created) {
		return customerrors.ErrPlayerCannotJoinTheGame
	}
	if len(game.players) == 50 {
		return customerrors.ErrPlayerCannotJoinTheGame
	}
	_, ok := game.players[player.Name()]
	if ok {
		return customerrors.ErrPlayerIsAdded
	}
	game.players[player.Name()] = player
	return nil
}

func (game *Game) onStart() error {
	if game.questions == nil {
		return customerrors.ErrGameHasNoQuestion
	}
	err := game.state.Started()
	if err == nil {
		game.currentQuestion = 1
	}
	return err
}

func (game *Game) startReceiveAnswer() error {
	answerChannel := make(chan question.AnswerFromPlayer)
	defer close(answerChannel)
	closeChannel := make(chan int8)
	for _, player := range game.players {
		go player.Listen(answerChannel, closeChannel)
	}
	currentQuestion, _ := game.CurrentQuestion()
	startTime := time.Now()
	answered := 0
	for {
		select {
		case playersAnswer := <-answerChannel:
			addPoint := game.calcAddPoint(
				currentQuestion.CompareAnswer(playersAnswer.Symbol()),
				time.Since(startTime).Milliseconds(),
			)
			game.leaderboard.AddPoint(
				playersAnswer.PlayerName(),
				addPoint,
			)
			answered += 1
			if answered == len(game.players) {
				close(closeChannel)
				return nil
			}
		case <-time.After(5 * time.Second):
			close(closeChannel)
			return customerrors.ErrQuestionTimeout
		}
	}
}

func (game *Game) calcAddPoint(
	isCorrect bool,
	answerDuration int64, //miliseconds
) uint32 {
	if !isCorrect {
		return 0
	}
	return rightAnswerMaxPoint - uint32(answerDuration)
}

func (game *Game) setQuestionList(
	input []question.Question,
) error {
	if !game.state.Is(Created) {
		return customerrors.ErrCannotSetListOfQuestion
	}
	game.questions = input
	return nil
}

func (game *Game) onNextQuestion() error {
	if game.currentQuestion == len(game.questions) {
		return customerrors.ErrCannotGoToTheNextQuestion
	}
	game.currentQuestion += 1
	return game.startReceiveAnswer()
}

func (game *Game) updateState(newState State) {
	game.state = newState
}
