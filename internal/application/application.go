package application

import (
	"amalia/internal/domain"
	"github.com/google/go-github/v61/github"
)

type ICodeRepositoryProvider interface {
	GetEventName() string
	GetRepository() string
	GetPullRequestChanges() ([]github.CommitFile, error)
	CreateComment(changedFiles []github.CommitFile) error
}

type App struct {
	codeRepositoryProvider ICodeRepositoryProvider
	aiModel                domain.IAIModel
}

func NewApp(
	codeRepositoryProvider ICodeRepositoryProvider,
	model domain.IAIModel) *App {
	return &App{
		codeRepositoryProvider: codeRepositoryProvider,
		aiModel:                model,
	}
}
