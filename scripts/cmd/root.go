package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "contentful-article-cli",
  Short: "Contentful article CLI is a CLI tool for local editing of contentful articles.",
  Run: func(cmd *cobra.Command, args []string) {
    // ルートコマンドはなにもしない
    fmt.Println("Execute any processing with the subcommand.")
  },
}

// Execute はコマンド全体の実行処理
func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
    cobra.OnInitialize()
    rootCmd.AddCommand(
        // サブコマンドを登録する
        pullCmd(),
        pushCmd(),
        showCmd(),
    )
}
