package domain

type IAIModel interface {
	GetComment(content string) (string, error)
}
