package jsonquestion

import (
	"errors"
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/hoanganh-ng/myhoot/helpers"
)

type NewRepositoryTestCase struct {
	QuestionPath string
	Expect       error
}

func TestNewRepositoryf(t *testing.T) {
	currentDir, _ := os.Getwd()
	testCases := []NewRepositoryTestCase{
		{
			// QuestionPath: "afnsl/djsks.json",
			QuestionPath: path.Join(currentDir, path.Dir("../../"), "afnsl", "djsks.json"),
			Expect:       fmt.Errorf("cannot open JSON questions file at %s", "afnsl/djsks.json"),
		},
		{
			QuestionPath: path.Join(currentDir, path.Dir("../../../"), "needle", "questions.json"),
			Expect:       nil,
		},
		{
			QuestionPath: "", //use default path
			Expect:       nil,
		},
	}
	for _, tc := range testCases {
		err := RunTestRepo(tc)
		if err != nil {
			t.Error(err)
		}
	}
}

func RunTestRepo(tc NewRepositoryTestCase) error {
	helpers.PrettyPrint(tc)
	_, err := NewRepository(tc.QuestionPath)
	if tc.Expect == nil && err != nil {
		return fmt.Errorf("expect no error. But got one|%w", err)
	}
	if tc.Expect != nil && err == nil {
		return errors.New("expect an error. But did not get")
	}
	return nil
}
