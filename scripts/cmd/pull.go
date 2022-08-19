package cmd

import (
	"fmt"
	"os"

	"contentful-article-cli/scripts/pull"

	"github.com/spf13/cobra"
)

// pullCmd はContentfulで管理している記事をローカルに取得する
func pullCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pull",
		Short: "Fetch articles locally from Contentful.",
		Run: func(cmd *cobra.Command, args []string) {
			var articleID string
			if len(args) >= 1 {
				articleID = args[0]
			}

			// 記事IDが指定されていない場合、全記事を取得する
			if articleID == "" {
				if err := pull.FetchArticles(); err != nil {
					os.Exit(1)
				}
				fmt.Println("pull successful. (articles)")
				return
			}

			// 記事IDが指定されている場合、指定された1件の記事を取得する
			if err := pull.FetchArticle(articleID); err != nil {
				os.Exit(1)
			}

			fmt.Println("pull successful. (article)")
		},
	}

	return cmd
}
