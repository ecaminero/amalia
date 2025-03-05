package github_connection

import (
	"errors"
	"fmt"
)

// Errores base
var (
	// Errores de configuración
	ErrMissingToken      = errors.New("GITHUB_TOKEN environment variable is not set")
	ErrMissingRepository = errors.New("GITHUB_REPOSITORY environment variable is not set")
	ErrInvalidRepoFormat = errors.New("invalid GITHUB_REPOSITORY format")
	ErrMissingPR         = errors.New("GITHUB_PR_NUMBER environment variable is not set")
	ErrInvalidPRNumber   = errors.New("invalid GITHUB_PR_NUMBER")
	ErrMissingEventName  = errors.New("GITHUB_EVENT_NAME environment variable is not set")
	ErrMissingCommitSHA  = errors.New("GITHUB_SHA environment variable is not set")

	// Errores de operaciones
	ErrListingPullRequest  = errors.New("error listing pull request files")
	ErrCreatingComment     = errors.New("error creating pull request comment")
	ErrFetchingPullRequest = errors.New("error fetching pull request")
)

// Funciones de ayuda para envolver errores
func WrapListingPRError(err error) error {
	return fmt.Errorf("%w: %v", ErrListingPullRequest, err)
}

func WrapCreatingCommentError(err error) error {
	return fmt.Errorf("%w: %v", ErrCreatingComment, err)
}

func WrapInvalidRepoFormatError(repoName string) error {
	return fmt.Errorf("%w: %s", ErrInvalidRepoFormat, repoName)
}

func WrapInvalidPRNumberError(err error) error {
	return fmt.Errorf("%w: %v", ErrInvalidPRNumber, err)
}
