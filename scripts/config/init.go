package config

import (
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const projectDirName = "contentful-article-cli-go"

// init はCLIの共通で行う初期化処理
func init() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))
	envFilePath := string(rootPath) + `/.env`

	err := godotenv.Load(envFilePath)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Load env error.")
		os.Exit(1)
	}
}

// LoadContentfulEnv はContentful SDKの接続情報を環境変数から読み込む
func LoadContentfulEnv() (string, string, error) {
	token := os.Getenv("CONTENTFUL_ACCESS_TOKEN")
	spaceID := os.Getenv("CONTENTFUL_SPACE_ID")

	if token == "" || spaceID == "" {
		m := fmt.Sprintf("Environment variable not set. [ token: %v, spaceID: %v ]", token, spaceID)
		fmt.Println(m)
		return "", "", errors.New(m)
	}

	return token, spaceID, nil
}

// LoadArticleDirEnv はローカルで記事を管理するディレクトリパスを環境変数から読み込む
func LoadArticleDirEnv() (string, error) {
	dirname := os.Getenv("ARTICLE_DIR")

	if dirname == "" {
		m := fmt.Sprintf("Environment variable not set. [ dirname: %v ]", dirname)
		fmt.Println(m)
		return "", errors.New(m)
	}

	return dirname, nil
}
