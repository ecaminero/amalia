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

	for _, change := range changes {
		println(change)
		err = a.aiModel.GetComment("change")
		if err != nil {
			log.Fatalf("Error: %v\n", err)
		}
	}

}
