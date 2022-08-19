package show

import (
	"contentful-article-cli/scripts/config"
	pkgtw "contentful-article-cli/scripts/pkg/tablewriter"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	dirname string
)

// initShow は初期化処理
func initShow() error {
	var err error
	dirname, err = config.LoadArticleDirEnv()
	if err != nil {
		return err
	}

	return nil
}

// ShowLocalArticles は`cmd/show`で実行されるメイン処理
func ShowLocalArticles() error {
	// 初期化する
	if err := initShow(); err != nil {
		return err
	}

	// ファイル名のリストを取得する
	filenames, err := filenames()
	if err != nil {
		return err
	}

	// ローカルで管理されている記事が存在しない場合、処理を終了する
	if filenames == nil {
		fmt.Println("There are no articles locally.")
		return nil
	}

	// ローカルで管理されている記事情報を表示する
	table := pkgtw.NewTable()
	for i, filename := range filenames {
		attrs := filenameToAttrs(filename)
		n := strconv.Itoa(i+1)
		attrs = append([]string{n}, attrs...)
		table.Append(attrs)
	}

	table.Render()

	return nil
}

// filenames はローカルで管理されている記事のファイル名をリストで取得する
func filenames() ([]string, error) {
	// ローカルで管理されている記事のファイルを、リストで取得する
	files, err := os.ReadDir(dirname)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// ファイルが0件だった場合、処理を終了する
	if len(files) == 0 {
		return nil, nil
	}

	// ローカルで管理されている記事のファイル名をリストにして返す
	var filenames []string
	for _, f := range files {
		filenames = append(filenames, f.Name())
	}

	return filenames, nil
}

// filenameToAttrs は指定されたファイル名から記事の要素の値をリストで取得する
func filenameToAttrs(filename string) []string {
	n := strings.TrimRight(filename, ".md")
	v := strings.Split(n, "--")
	attrs := []string{v[0], v[2], v[1]}

	return attrs
}
