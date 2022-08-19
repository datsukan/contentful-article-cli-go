package push

import (
	pkgcf "contentful-article-cli/scripts/pkg/contentful"
	"fmt"
)

// initValidate は初期化処理
func initValidate() error {
	var err error
	cma, space, err = pkgcf.NewContentfulSDK()
	if err != nil {
		return err
	}

	return nil
}

// Valid は`push`パッケージで利用するパラメータ検証処理
func Valid(args []string) (bool, error) {
	// パラメータが0件の場合、不正とする
	if len(args) == 0 {
		fmt.Println("Article ID is not specified in the first argument.")
		return false, nil
	}

	articleID := args[0]
	exist, err := exist(articleID)
	if err != nil {
		return false, err
	}

	// パラメータで指定された記事がContentfulに存在しない場合、不正とする
	if !exist {
		fmt.Println("Article not found.")
		return false, nil
	}

	return true, nil
}

// exist は指定された記事IDがContentfulに存在することを判定する
func exist(articleID string) (bool, error) {
	// 初期化する
	if err := initValidate(); err != nil {
		return false, err
	}

	// Contentfulから記事情報を取得する
	entry, err := cma.Entries.Get(space.Sys.ID, articleID)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	if entry == nil {
		return false, nil
	}

	return true, nil
}
