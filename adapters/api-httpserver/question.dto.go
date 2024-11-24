package apihttpserver

type CreateQuestionRequest struct {
	Question     string   `json:"question" validate:"required|min_len:1" label:"Question"`
	WrongAnswers []string `json:"wrong_answers" validate:"required|len:3" label:"Wrong Answers"`
	RightAnswer  string   `json:"right_answer" validate:"required|min_len:1" label:"Right Answers"`
}
