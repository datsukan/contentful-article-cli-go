package tablewriter

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

// NewTable はtablewriterのインスタンスを生成する
func NewTable() *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"#", "id", "slug", "title"})
	return table
}

