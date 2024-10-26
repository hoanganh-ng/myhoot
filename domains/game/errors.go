package game

import "errors"

var (
	ErrInvalidQuestion           = errors.New("invalid question")
	ErrCannotSetListOfQuestion   = errors.New("can not set list of question")
	ErrCannotGoToTheNextQuestion = errors.New("can not go to next question")
	ErrGameIsNotStartYet         = errors.New("the game is not running")
	ErrPlayerCannotJoinTheGame   = errors.New("cannot join the game")
	ErrPlayerIsAdded             = errors.New("player've already been added")
	ErrQuestionTimeout           = errors.New("unable to answer the question")
	ErrPlayerIsNotExisted        = errors.New("player is not existed")
)
