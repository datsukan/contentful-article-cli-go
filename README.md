# Contentful article CLI by Golang

This is a CLI tool for writing article text managed by Contentful in Markdown in Visual Studio Code.

## 🚀 Quick start

```
git clone https://github.com/datsukan/contentful-article-cli-go.git
cd contentful-article-cli-go
```

### Install

- [markdownlint](https://marketplace.visualstudio.com/items?itemName=DavidAnson.vscode-markdownlint) (Optional)
- [Markdown All in One](https://marketplace.visualstudio.com/items?itemName=yzhang.markdown-all-in-one) (Optional)

### Copy `.env`

```
cp .env.example .env
```

### Setting `.env`

```
CONTENTFUL_SPACE_ID=xxxxxxxxxx
CONTENTFUL_ACCESS_TOKEN=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

### Setting Contentful

make a `content model` and `entry`.

#### Content model Fields

Article content model

- slug (Short text)
- title (Short text)
- body (Long text)

## ❓ How to get

### Space id

[ Contentful dashboard > Settings > General settings > Space ID ]

### Access token

[ Contentful dashboard > Settings > API Keys > Content management tokens > Generate personal token ]

### Article id

[ Contentful dashboard > Content > entry item > Sidebar > info > ENTRY ID ]

## 🧐 Usage

```sh
./contentful-article-cli help
```

```
Contentful article CLI is a CLI tool for local editing of contentful articles.

Usage:
  contentful-article-cli [flags]
  contentful-article-cli [command]

Available Commands:
  help        Help about any command
  pull        Fetch articles locally from Contentful.
  push        Update Contentful with local articles.
  show        Display information for local articles.

Flags:
  -h, --help   help for contentful-article-cli

Use "contentful-article-cli [command] --help" for more information about a command.
```

## 📝 Markdown rule

`.markdownlint.jsonc` is the markdown rule configuration file.

### Reference page

- [rules](https://github.com/DavidAnson/markdownlint/blob/main/doc/Rules.md)
- [manual](https://marketplace.visualstudio.com/items?itemName=DavidAnson.vscode-markdownlint#configure)
