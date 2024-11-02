package game

import "fmt"

type GameService struct {
	gameRepo     GameRepository
	questionRepo QuestionRepository
}

const (
	DEFAULT_NUMBER_OF_QUESTION_PER_GAME int = 20
)

func (gs GameService) CreateGame(
	manager *Manager,
) error {
	err := manager.CreateGame()
	if err != nil {
		return fmt.Errorf("manager.CreateGame|%w", err)
	}
	return nil
}

func (gs GameService) RandomQuestion(
	manager *Manager,
	numberOfQuestion int,
) error {
	listOfQuestion, err := gs.questionRepo.Random(numberOfQuestion)
	if err != nil {
		return fmt.Errorf("questionRepo.Random|%w", err)
	}
	err = manager.SelectQuestionList(listOfQuestion)
	if err != nil {
		return fmt.Errorf("manager.SelectQuestionList|%w", err)
	}
	return err
}

func (gs GameService) JoinGame(
	player *Player,
	gameID string,
	submittedPWD string,
) error {
	existedGame, err := gs.gameRepo.Get(gameID)
	if err != nil {
		return fmt.Errorf("gameRepo.Get|%w", err)
	}
	correctedPWD := existedGame.Authenticate(submittedPWD)
	if !correctedPWD {
		return ErrInvalidGamePwd
	}
	err = existedGame.OnPlayerRequestJoin(player)
	if err != nil {
		return fmt.Errorf("OnPlayerRequestJoin|%w", err)
	}
	return err
}
