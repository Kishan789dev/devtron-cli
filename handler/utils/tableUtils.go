package utils

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func FormatTable(table table.Table) table.Table {
	headerFmt := color.New(color.FgBlue, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	table.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	return table
}
