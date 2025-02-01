package jsonquestion

type Answer struct {
	Value     string `json:"value"`
	IsCorrect bool   `json:"is_correct"`
}

type Question struct {
	Value   string   `json:"value"`
	Answers []Answer `json:"answers"`
}
