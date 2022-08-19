package cmd

import (
	"fmt"
	"os"

	"contentful-article-cli/scripts/push"

	"github.com/spf13/cobra"
)

// pushCmd はローカルで管理している記事をContentfulへ反映する
func pushCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "push",
		Short: "Update Contentful with local articles.",
		Run: func(cmd *cobra.Command, args []string) {
			valid, err := push.Valid(args)
			if err != nil {
				os.Exit(1)
			}

			// パラメーターが不正な場合は処理を終了する
			if !valid {
				fmt.Println("validation error")
				return
			}

			articleID := args[0]
			if err := push.UpdateArticle(articleID); err != nil {
				os.Exit(1)
			}

			fmt.Println("push successful.")
		},
	}

	return cmd
}
