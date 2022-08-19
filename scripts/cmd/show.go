package cmd

import (
	"fmt"
	"os"

	"contentful-article-cli/scripts/show"

	"github.com/spf13/cobra"
)

// showCmd はローカルで管理している記事の情報を表示する
func showCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "Display information for local articles.",
		Run: func(cmd *cobra.Command, args []string) {
			if err := show.ShowLocalArticles(); err != nil {
				os.Exit(1)
			}

			fmt.Println("show successful.")
		},
	}

	return cmd
}
