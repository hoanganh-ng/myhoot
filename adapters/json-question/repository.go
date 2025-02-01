package jsonquestion

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path"

	"github.com/hoanganh-ng/myhoot/helpers"
)

var (
	currentDir, _         = os.Getwd()
	default_question_path = path.Join(currentDir, path.Dir("../../../"), "needle", "questions.json")
)

type Repository struct {
	inMemory []Question
}

func NewRepository(questionPath string) (*Repository, error) {
	if questionPath == "" {
		questionPath = default_question_path
	}
	repo := &Repository{
		inMemory: make([]Question, 0),
	}
	err := repo.load(questionPath)
	return repo, err
}

func (repo *Repository) load(strPath string) error {
	data, err := os.ReadFile(strPath)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("cannot open JSON questions file at %s", strPath)
	}
	err = json.Unmarshal(data, &repo.inMemory)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("cannot decode JSON questions file at %s", strPath)
	}
	return nil
}

func (repo Repository) Create(*Question) error {
	return nil
}

func (repo Repository) Random(n int) ([]Question, error) {
	if len(repo.inMemory) > n {
		return nil, errors.New("not enough question for the request")
	}
	var (
		checkM = make(map[int]bool)
		result = make([]Question, 0)
		enough = false
	)

	for enough {
		rInt := rand.New(helpers.RandSource).Intn(len(repo.inMemory))
		_, ok := checkM[rInt]
		if !ok {
			checkM[rInt] = true
			result = append(result, repo.inMemory[rInt])
			if len(result) == n {
				break
			}
		}
	}
	return result, nil
}
