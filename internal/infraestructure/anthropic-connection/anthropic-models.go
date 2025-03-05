package anthropic_connection

import (
	"github.com/anthropics/anthropic-sdk-go"
	"strings"
)

var ModelMap = map[string]anthropic.Model{
	"3-5-sonnet": anthropic.ModelClaude3_5SonnetLatest,
	"3.7-sonnet": anthropic.ModelClaude3_7SonnetLatest,
	"opus":       anthropic.ModelClaude3OpusLatest,
	"haiku":      anthropic.ModelClaude3_5HaikuLatest,
}

func GetModel(key string) (anthropic.Model, bool) {
	normalizedKey := strings.ToLower(strings.TrimSpace(key))
	model, found := ModelMap[normalizedKey]
	return model, found
}
