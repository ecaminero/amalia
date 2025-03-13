package github_connection

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-github/v61/github"
)

type GithubConnection struct {
	Client            *github.Client `json:"client"`
	RepositoryName    string         `json:"repository_name"`
	RepoOwner         string         `json:"repo_owner,omitempty"`
	PullRequestNumber int            `json:"pull_request_number"`
	EventName         string         `json:"event_name"`
	GetCommitSHA      string         `json:"get_commit_sha"`
}

type FileChange struct {
	github.CommitFile
	Position int `json:"position"`
}

func NewGithubConnection() (*GithubConnection, error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return nil, ErrMissingToken
	}

	repoFullName := os.Getenv("GITHUB_REPOSITORY")
	if repoFullName == "" {
		return nil, ErrMissingRepository
	}
	repoParts := strings.Split(repoFullName, "/")
	if len(repoParts) != 2 {
		return nil, WrapInvalidRepoFormatError(repoFullName)
	}
	repoOwner := repoParts[0]
	githubRepositoryName := repoParts[1]

	pullRequestNumberStr := os.Getenv("GITHUB_PR_NUMBER")
	if pullRequestNumberStr == "" {
		return nil, ErrMissingPR
	}
	pullRequestNumber, err := strconv.Atoi(pullRequestNumberStr)
	if err != nil {
		return nil, WrapInvalidPRNumberError(err)
	}

	eventName := os.Getenv("GITHUB_EVENT_NAME")
	if eventName == "" {
		return nil, ErrMissingEventName
	}

	githubCommitSHA := os.Getenv("GITHUB_SHA")
	if githubCommitSHA == "" {
		return nil, ErrMissingCommitSHA
	}

	GithubClient := github.NewClient(nil).WithAuthToken(token)

	return &GithubConnection{
		Client:            GithubClient,
		RepositoryName:    githubRepositoryName,
		RepoOwner:         repoOwner,
		PullRequestNumber: pullRequestNumber,
		EventName:         eventName,
		GetCommitSHA:      githubCommitSHA,
	}, nil
}

func (receiver *GithubConnection) GetEventName() string {
	return receiver.EventName
}

func (receiver *GithubConnection) GetRepository() string {
	return receiver.RepositoryName
}

func (receiver *GithubConnection) GetPullRequestChanges() ([]github.CommitFile, error) {
	var ctx = context.Background()
	var opt = &github.ListOptions{PerPage: 100}
	var allChanges []github.CommitFile

	for {
		files, resp, err := receiver.Client.PullRequests.ListFiles(
			ctx, receiver.RepoOwner,
			receiver.RepositoryName,
			receiver.PullRequestNumber,
			opt)

		if err != nil {
			return nil, WrapListingPRError(err)
		}
		for _, file := range files {
			if file.GetAdditions() == 0 {
				continue
			}
			allChanges = append(allChanges, *file)
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return allChanges, nil
}

func (receiver *GithubConnection) CreateComment(files []github.CommitFile) error {
	ctx := context.Background()

	for _, file := range files {
		comments := analyzeFileAndCreateComments(&file)
		for _, commentData := range comments {
			_, _, err := receiver.Client.PullRequests.CreateComment(
				ctx,
				receiver.RepoOwner,
				receiver.RepositoryName,
				receiver.PullRequestNumber,
				commentData)
			if err != nil {
				return WrapCreatingCommentError(err)
			}
		}
	}

	return nil
}

func analyzeFileAndCreateComments(file *github.CommitFile) []*github.PullRequestComment {
	// El resto del código se mantiene igual
	var comments []*github.PullRequestComment
	patch := file.GetPatch()
	lines := strings.Split(patch, "\n")

	position := 0
	newLineNumber := 0
	startLine := 0
	endLine := 0
	var newLines []string

	for _, line := range lines {
		position++
		if strings.HasPrefix(line, "+") {
			newLineNumber++
			if startLine == 0 {
				startLine = newLineNumber
			}
			endLine = newLineNumber
			newLines = append(newLines, strings.TrimPrefix(line, "+"))
		} else {
			if startLine != 0 {
				comment := createCommentForLines(file, position-1, startLine, endLine, newLines)
				comments = append(comments, comment)
				startLine = 0
				endLine = 0
				newLines = nil
			}
		}
	}
	// Handle case where file ends with new lines
	if startLine != 0 {
		comment := createCommentForLines(file, position, startLine, endLine, newLines)
		comments = append(comments, comment)
	}
	return comments
}

func createCommentForLines(file *github.CommitFile, position, startLine, endLine int, lines []string) *github.PullRequestComment {
	commentBody := fmt.Sprintf("Code Review %s - New lines %d to %d:\n\n", time.Now().Format("2006-01-02 15:04:05"), startLine, endLine)
	commentBody += strings.Join(lines, "\n")

	return &github.PullRequestComment{
		Body:     github.String(commentBody),
		CommitID: github.String(os.Getenv("GITHUB_SHA")),
		Path:     github.String(file.GetFilename()),
		Position: github.Int(position),
	}
}
