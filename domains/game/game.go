package game

import (
	"time"

	"github.com/google/uuid"
)

const rightAnswerMaxPoint uint32 = 5 * 1000 //max is reached when player anwser in 0 miliseconds

type Participant interface {
	ReceiveLeaderboard()
}

type Game struct {
	id              uuid.UUID
	state           State
	manager         Manager
	currentQuestion int
	players         map[string]*Player
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
	answerChannel := make(chan AnswerFromPlayer)
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
				currentQuestion.RightAnser().Symbol == playersAnswer.symbol,
				time.Since(startTime).Milliseconds(),
			)
			game.leaderboard.AddPoint(
				playersAnswer.playerName,
				addPoint,
			)
			answered += 1
			if answered == len(game.players) {
				close(closeChannel)
				return nil
			}
		case <-time.After(5 * time.Second):
			close(closeChannel)
			return ErrQuestionTimeout
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

func (game *Game) PlayerJoin(player *Player) error {
	if !game.state.Is(Created) {
		return ErrPlayerCannotJoinTheGame
	}
	if len(game.players) == 50 {
		return ErrPlayerCannotJoinTheGame
	}
	_, ok := game.players[player.Name()]
	if ok {
		return ErrPlayerIsAdded
	}
	game.players[player.Name()] = player
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
	return game.startReceiveAnswer()
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
