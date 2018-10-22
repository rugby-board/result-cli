package cmd

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/rugby-board/result-cli/match"
)

const maxColWidth = 120

// OutputMarkdownTable print a formatted markdown table
func OutputMarkdownTable(matches []*match.Match) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"主队", "比分", "客队"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetColWidth(maxColWidth)
	table.SetCenterSeparator("|")
	for _, m := range matches {
		score := fmt.Sprintf("%d-%d", m.Team1Score, m.Team2Score)
		table.Append([]string{m.Team1Name, score, m.Team2Name})
	}
	table.Render()
}
