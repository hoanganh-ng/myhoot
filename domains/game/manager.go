package game

import "github.com/hoanganhnguyen17/myhoot/domains/user"

type Manager struct {
	*user.User
	game *Game
}
