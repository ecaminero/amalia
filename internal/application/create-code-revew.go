package application

import (
	"log"
)

type Context struct {
	Owner             string
	Repository        string
	PullRequestNumber int
	Token             string
}

func (a *App) CreateCodeReview() {
	changes, err := a.codeRepositoryProvider.GetPullRequestChanges()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	err = a.aiModel.GetComment(changes)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

}
