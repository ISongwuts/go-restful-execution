package usecase

import (
	"os"
	"os/exec"
	"strings"
	"github.com/ISongwuts/go-restful-execution/internal/mock"
	"github.com/ISongwuts/go-restful-execution/packages/domain"
)

type(
	Submit domain.Submit
	Factory interface {
		GetLanguage() string
		GetCode() string
		GetSubmitAt() string
		GetStatus() string
		TestCase() (string, string, error)
	}
)

func NewSubmit(lang, code, submitAt, status string) Factory {
	return &Submit{
		Language: lang,
		Code: code,
		SubmitAt: submitAt,
		Status: status,
	}
}

func (s *Submit) GetLanguage() string {
	return s.Language
}

func (s *Submit) GetCode() string {
	
	return s.Code
}

func (s *Submit) GetSubmitAt() string {
	return s.SubmitAt
}

func (s *Submit) GetStatus() string {
	return s.Status
}

func (s *Submit) TestCase() (string, string, error) {
	cwd, _ := os.Getwd()
	path := cwd + "/internal/languages/js/test.js"
	err := os.WriteFile(path, []byte(s.Code), 0644)
	if err != nil {
		return "", "", err
	}

	cmd := exec.Command("node", path)
	raw_output, err := cmd.CombinedOutput()

	if err != nil {
		return "", "", err
	}

	output := strings.TrimRight(string(raw_output), "\n")
	status := "not pass"

	if output == mock.OUTPUT_MOCK_1 {
		status = "pass"
	}
	
	return output, status, nil
}