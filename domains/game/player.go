package game

import "github.com/hoanganhnguyen17/myhoot/domains/user"

type Player struct {
	user.User
	game *Game
}

func (p Player) InGame() bool {
	return p.game != nil
}

func (p Player) SubmitAnswer(
	question Question,
	answer Answer,
) {
	p.game.ReceiveAnswer()
}
