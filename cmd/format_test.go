package cmd

import (
	"testing"

	"github.com/rugby-board/rugby-result/match"
)

func TestMarkdownTable(t *testing.T) {
	matches := make([]*match.Match, 1)
	matches[0] = &match.Match{
		Team1Name:  "South Africa",
		Team1Score: 23,
		Team2Name:  "England",
		Team2Score: 12,
	}
	markdownString := OutputMarkdownTable(matches)
	if markdownString != `|     主队     | 比分  |  客队   |
|--------------|-------|---------|
| South Africa | 23-12 | England |
` {
		t.Error("Output error")
	}
}
