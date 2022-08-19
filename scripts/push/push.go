package push

import (
	"contentful-article-cli/scripts/config"
	pkgaf "contentful-article-cli/scripts/pkg/articlefile"
	pkgcf "contentful-article-cli/scripts/pkg/contentful"
	pkgrp "contentful-article-cli/scripts/pkg/resparse"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/contentful-labs/contentful-go"
)

var (
	cma *contentful.Client
	space *contentful.Space
	dirname string
)

// initPush は初期化処理
func initPush() error {
	var err error

	// Contentful SDK のクライアントインスタンスを生成する
	cma, space, err = pkgcf.NewContentfulSDK()
	if err != nil {
		return err
	}

	// ローカルで記事を管理するディレクトリを環境変数から取得する
	dirname, err = config.LoadArticleDirEnv()
	if err != nil {
		return err
	}

	return nil
}

// UpdateArticle は`cmd/push`で実行されるメイン処理
func UpdateArticle(articleID string) error {
	// 初期化する
	if err := initPush(); err != nil {
		return err
	}

	// Contentfulから記事情報を取得する
	entry, err := cma.Entries.Get(space.Sys.ID, articleID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Contentfulから記事情報が取得できない場合、処理を終了する
	if entry == nil {
		fmt.Println("Article not found.")
		return errors.New("article not found")
	}

	// ローカルで管理されている記事の本文を取得する
	text, err := fileText(entry)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Contentfulから取得した記事情報に、ローカルで管理されている記事の本文を反映する
	body, err := pkgrp.StringToField(text)
	if err != nil {
		return err
	}
	entry.Fields["body"] = body

	// ローカルで管理されている記事の本文を、Contentfulへ反映する
	if err := cma.Entries.Upsert(space.Sys.ID, entry); err != nil {
		return err
	}

	return nil
}

// fileText はローカルで管理されている記事の本文を取得する
func fileText(entry *contentful.Entry) (string, error) {
	// ローカルで管理されている記事のファイル名を取得する
	filename, err := pkgaf.ArticleToFilename(entry)
	if err != nil {
		return "", err
	}

	pathname := fmt.Sprintf("%v/%v", dirname, filename)

	// ローカルで管理している記事の本文を取得する
	bytes, err := ioutil.ReadFile(pathname)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
