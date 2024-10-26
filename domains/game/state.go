package game

import "errors"

// type State int

const (
	Created = iota
	Started
	Paused
	Stopped
	Completed
)

var ErrInvalidStateChanging = errors.New("invalid state changing")

type State interface {
	Created() error
	Started() error
	Paused() error
	Stopped() error
	Completed() error
	Value() int
	Is(int) bool
}

//
type CreatedState struct {
	game *Game
}

func (s CreatedState) Created() error { return ErrInvalidStateChanging }
func (s CreatedState) Started() error {
	s.game.updateState(StartedState{game: s.game})
	return nil
}
func (s CreatedState) Paused() error { return ErrInvalidStateChanging }
func (s CreatedState) Stopped() error {
	s.game.updateState(StoppedState{game: s.game})
	return nil
}
func (s CreatedState) Completed() error { return ErrInvalidStateChanging }
func (s CreatedState) Value() int       { return Created }
func (s CreatedState) Is(i int) bool    { return Created == i }

//
type StartedState struct {
	game *Game
}

func (s StartedState) Created() error { return ErrInvalidStateChanging }
func (s StartedState) Started() error { return ErrInvalidStateChanging }
func (s StartedState) Paused() error {
	s.game.updateState(PausedState{game: s.game})
	return nil
}
func (s StartedState) Stopped() error {
	s.game.updateState(StoppedState{game: s.game})
	return nil
}
func (s StartedState) Completed() error {
	s.game.updateState(CompletedState{game: s.game})
	return nil
}
func (s StartedState) Value() int    { return Started }
func (s StartedState) Is(i int) bool { return Started == i }

//
type PausedState struct {
	game *Game
}

func (s PausedState) Created() error { return ErrInvalidStateChanging }
func (s PausedState) Started() error {
	s.game.updateState(StartedState{game: s.game})
	return nil
}
func (s PausedState) Paused() error { return ErrInvalidStateChanging }
func (s PausedState) Stopped() error {
	s.game.updateState(StoppedState{game: s.game})
	return nil
}
func (s PausedState) Completed() error { return ErrInvalidStateChanging }
func (s PausedState) Value() int       { return Paused }
func (s PausedState) Is(i int) bool    { return Paused == i }

//
type StoppedState struct {
	game *Game
}

func (s StoppedState) Created() error   { return ErrInvalidStateChanging }
func (s StoppedState) Started() error   { return ErrInvalidStateChanging }
func (s StoppedState) Paused() error    { return ErrInvalidStateChanging }
func (s StoppedState) Stopped() error   { return ErrInvalidStateChanging }
func (s StoppedState) Completed() error { return ErrInvalidStateChanging }
func (s StoppedState) Value() int       { return Stopped }
func (s StoppedState) Is(i int) bool    { return Stopped == i }

//
type CompletedState struct {
	game *Game
}

func (s CompletedState) Created() error   { return ErrInvalidStateChanging }
func (s CompletedState) Started() error   { return ErrInvalidStateChanging }
func (s CompletedState) Paused() error    { return ErrInvalidStateChanging }
func (s CompletedState) Stopped() error   { return ErrInvalidStateChanging }
func (s CompletedState) Completed() error { return ErrInvalidStateChanging }
func (s CompletedState) Value() int       { return Completed }
func (s CompletedState) Is(i int) bool    { return Completed == i }
