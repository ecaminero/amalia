package main

import (
	github_connection "amalia/internal/infraestructure/github-connection"
	stub_persistence "amalia/internal/infraestructure/stub"
	"amalia/internal/application"

	"fmt"
	"log"
	"os"
	"strings"
)

func print_all_variables() {
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if strings.HasPrefix(pair[0], "GITHUB_") {
			fmt.Println(pair[0], ":", pair[1])
		}
	}
}

func main() {
	githubConnection, err := github_connection.NewGithubConnection()
	if err != nil {
		log.Fatal(err)
	}

	aiModel := stub_persistence.NewStubModelRepository()
	app := application.NewApp(githubConnection, aiModel)
	eventName := githubConnection.GetEventName()
	switch eventName {
	case "pull_request_target", "pull_request":
		app.CreateCodeReview()
	case "pull_request_review_comment":
		// application.HandleCommentReview(repoOwner, githubRepositoryName, prNumber)
		print("HandleCommentReview")
	default:
		log.Fatalf("Skipped: current event is %s, only support pull_request event", eventName)
	}
}
