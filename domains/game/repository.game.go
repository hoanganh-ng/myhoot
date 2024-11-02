package game

type GameRepository interface {
	Get(gameID string) (*Game, error)
	Persist(*Game) error
}
