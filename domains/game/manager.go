package game

import (
	"github.com/hoanganh-ng/myhoot/domains/participant"
	"github.com/hoanganh-ng/myhoot/domains/question"
	customerrors "github.com/hoanganh-ng/myhoot/packages/custom-errors"
)

type Manager struct {
	*participant.User
	game *Game
}

func (manager *Manager) CreateGame() error {
	if manager.game != nil {
		return customerrors.ErrAGameIsNotRelease
	}
	nGame := newGame()
	manager.game = nGame
	return nil
}

func (manager *Manager) SelectQuestionList(
	listOfQuestion []question.Question,
) error {
	return manager.game.setQuestionList(listOfQuestion)
}

func (manager Manager) StartGame() error {
	return manager.game.onStart()
}

func (manager Manager) NextQuestion() error {
	return manager.game.onNextQuestion()
}
