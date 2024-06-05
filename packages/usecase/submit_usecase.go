package usecase

import "github.com/ISongwuts/go-restful-execution/packages/domain"

type(
	Submit domain.Submit
	Factory interface {
		NewSubmit(lang, code, submitAt, status string) Factory
	}
)

func (s *Submit) NewSubmit(lang, code, submitAt, status string) Factory {
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