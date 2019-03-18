package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/rugby-board/rugby-result/match"
)

// Reference: https://github.com/kami-zh/go-capturer/blob/master/main.go
func captureOutput(f func()) string {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout := os.Stdout
	os.Stdout = w
	defer func() {
		os.Stdout = stdout
	}()

	f()
	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestMarkdownTable(t *testing.T) {
	matches := make([]*match.Match, 1)
	matches[0] = &match.Match{
		Team1Name:  "South Africa",
		Team1Score: 23,
		Team2Name:  "England",
		Team2Score: 12,
	}
	markdownString := captureOutput(func() {
		OutputMarkdownTable(matches)
	})
	fmt.Println(markdownString)
	if markdownString != `|     主队     | 比分  |  客队   |
|--------------|-------|---------|
| South Africa | 23-12 | England |
` {
		t.Error("Output error")
	}
}
