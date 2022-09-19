package articlefile

import (
	"testing"

	"github.com/contentful-labs/contentful-go"
	"github.com/stretchr/testify/assert"
)

func TestArticleAttr(t *testing.T) {
	tests := []struct {
		name string
		entry *contentful.Entry
		wantSlug string
		wantTitle string
		wantBody string
	}{
		{
			name: "記事情報が正しく取得できること",
			entry: &contentful.Entry{
				Fields: map[string]interface{}{
					"slug": map[string]interface{}{"ja": "test-1"},
					"title": map[string]interface{}{"ja": "test title"},
					"body": map[string]interface{}{"ja": "test article body."},
				},
			},
			wantSlug: "test-1",
			wantTitle: "test title",
			wantBody: "test article body.",
		},
		{
			name: "記事のタイトルに'/'が含まれる場合、全て'／'に置換されて取得されること",
			entry: &contentful.Entry{
				Fields: map[string]interface{}{
					"slug": map[string]interface{}{"ja": "test-1"},
					"title": map[string]interface{}{"ja": "test/title/slash"},
					"body": map[string]interface{}{"ja": "test article body."},
				},
			},
			wantSlug: "test-1",
			wantTitle: "test／title／slash",
			wantBody: "test article body.",
		},
		{
			name: "Contentfulから取得したデータに対象の項目が含まれない場合、空文字で取得されること",
			entry: &contentful.Entry{
				Fields: map[string]interface{}{
					"hoge": map[string]interface{}{"ja": "hoge"},
				},
			},
			wantSlug: "",
			wantTitle: "",
			wantBody: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slug, title, body, err := ArticleAttr(tt.entry)

			assert.NoError(t, err)
			assert.Equal(t, tt.wantSlug, slug)
			assert.Equal(t, tt.wantTitle, title)
			assert.Equal(t, tt.wantBody, body)
		})
	}
}

func TestAttrToFilename(t *testing.T) {
	id := "test-id"
	title := "test title"
	slug := "test-slug"
	want := "test-id--test title--test-slug.md"

	filename := AttrToFilename(id, title, slug)

	assert.Equal(t, want, filename, "想定通りのファイル名が取得できること")
}

func TestArticleToFilename(t *testing.T) {
	entry := &contentful.Entry{
		Sys: &contentful.Sys{
			ID: "test-id",
		},
		Fields: map[string]interface{}{
			"slug": map[string]interface{}{"ja": "test-slug"},
			"title": map[string]interface{}{"ja": "test title"},
			"body": map[string]interface{}{"ja": "test article body."},
		},
	}

	want := "test-id--test title--test-slug.md"

	filename, err := ArticleToFilename(entry)

	assert.NoError(t, err, "エラーが発生していないこと")
	assert.Equal(t, want, filename)
}
