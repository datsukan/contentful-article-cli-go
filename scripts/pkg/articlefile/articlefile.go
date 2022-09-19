package articlefile

import (
	"fmt"
	"strings"

	pkgrp "contentful-article-cli/scripts/pkg/resparse"

	"github.com/contentful-labs/contentful-go"
)

// ArticleAttr は指定された記事情報から主要な項目の値を取得する
func ArticleAttr(entry *contentful.Entry) (string, string, string, error) {
	var slug, title, body string
	var err error
	for attr, field := range entry.Fields {
		switch attr {
		case "slug":
			slug, err = pkgrp.FieldToString(field)
			if err != nil {
				fmt.Println(err)
				return "", "", "", err
			}
		case "title":
			value, err := pkgrp.FieldToString(field)
			if err != nil {
				fmt.Println(err)
				return "", "", "", err
			}
			title = strings.Replace(value, "/", "／", -1)
		case "body":
			body, err = pkgrp.FieldToString(field)
			if err != nil {
				fmt.Println(err)
				return "", "", "", err
			}
		}
	}

	return slug, title, body, nil
}

// AttrToFilename は指定された記事の主要な情報をファイル名に変換する
func AttrToFilename(id, title, slug string) string {
	return fmt.Sprintf("%v--%v--%v.md", id, title, slug)
}

// ArticleToFilename は指定された記事をファイル名に変換する
func ArticleToFilename(entry *contentful.Entry) (string, error) {
	slug, title, _, err := ArticleAttr(entry)
	if err != nil {
		return "", err
	}

	return AttrToFilename(entry.Sys.ID, title, slug), nil
}
