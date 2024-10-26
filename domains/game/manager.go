package game

import "github.com/hoanganhnguyen17/myhoot/domains/participant"

type Manager struct {
	*participant.User
	game *Game
}

func (manager *Manager) CreateGame() error {
	newGame := New()
	manager.game = newGame
	return nil
}

func (manager *Manager) SelectQuestionList(
	listOfQuestion []Question,
) error {
	return manager.game.setQuestionList(listOfQuestion)
}

func (manager Manager) StartGame() error {
	return manager.game.onStart()
}

func (manager Manager) NextQuestion() error {
	return manager.game.onNextQuestion()
}
