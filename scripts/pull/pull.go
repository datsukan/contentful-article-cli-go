package pull

import (
	"fmt"
	"os"
	"strconv"

	"contentful-article-cli/scripts/config"
	pkgaf "contentful-article-cli/scripts/pkg/articlefile"
	pkgcf "contentful-article-cli/scripts/pkg/contentful"
	pkgtw "contentful-article-cli/scripts/pkg/tablewriter"

	contentful "github.com/contentful-labs/contentful-go"
	"github.com/olekukonko/tablewriter"
)

var (
	cma *contentful.Client
	space *contentful.Space
	dirname string
)

// initPull は初期化処理
func initPull() error {
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

	// ローカルで管理している記事をディレクトリごと削除する
	os.RemoveAll(dirname)

	// ローカルで記事を管理するディレクトリを作成する
	if err := os.Mkdir(dirname, 0755); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// FetchArticle は`cmd/pull`で実行されるメイン処理
// 指定された1件の記事を取得する
func FetchArticle(articleID string) error {
	// 初期化する
	if err := initPull(); err != nil {
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
		return nil
	}

	// 記事情報をローカルに保存して、保存対象の記事情報を表示する
	table := pkgtw.NewTable()
	save(entry, table)
	table.Render()

	return nil
}

// FetchArticle は`cmd/pull`で実行されるメイン処理
// 全件の記事を取得する
func FetchArticles() error {
	// 初期化する
	if err := initPull(); err != nil {
		return err
	}

	//  Contentfulから記事情報のリストを取得する
	collection := cma.Entries.List(space.Sys.ID)
	collection, err := collection.Next()
	if err != nil {
		fmt.Println(err)
		return err
	}

	// 記事情報をローカルに保存して、保存対象の記事情報を表示する
	table := pkgtw.NewTable()
	for _, entry := range collection.ToEntry() {
		save(entry, table)
	}

	table.Render()

	return nil
}

// createFile は指定されたファイル名とデータで新規にファイルを作成する
func createFile(name string, data []byte) error {
	if err := os.WriteFile(name, data, 0644); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// save は指定された記事情報をローカルに保存して、保存対象の記事情報を表示する
func save(entry *contentful.Entry, table *tablewriter.Table) error {
	// 記事情報から主要な項目の値を取得する
	slug, title, body, err := pkgaf.ArticleAttr(entry)
	if err != nil {
		return err
	}

	// Slug または タイトルが取得できない場合、処理を終了する
	if slug == "" || title == "" {
		return nil
	}

	// 表示する記事情報の一覧に追加する
	table.Append([]string{strconv.Itoa(table.NumLines()+1), entry.Sys.ID, slug, title})

	// ローカルにファイルを新規作成して、記事情報を保存する
	filename := pkgaf.AttrToFilename(entry.Sys.ID, title, slug)
	pathname := fmt.Sprintf("%v/%v", dirname, filename)
	createFile(pathname, []byte(body))

	return nil
}
