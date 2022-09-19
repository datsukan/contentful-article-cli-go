package resparse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldToString(t *testing.T) {
	tests := []struct {
		name  string
		field map[string]interface{}
		want  string
	}{
		{
			name:  "想定通りの値に変換できること",
			field: map[string]interface{}{"ja": "test field value"},
			want:  "test field value",
		},
		{
			name:  "キーが不正の場合は値が取得できないこと",
			field: map[string]interface{}{"en": "test field value"},
			want:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := FieldToString(tt.field)

			assert.NoError(t, err, "エラーが発生していないこと")
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestStringToField(t *testing.T) {
	tests := []struct {
		name string
		text string
		want map[string]interface{}
	}{
		{
			name: "想定通りの値に変換できること",
			text: "test field value",
			want: map[string]interface{}{"ja": "test field value"},
		},
		{
			name: "空文字でも想定通りの値に変換できること",
			text: "",
			want: map[string]interface{}{"ja": ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringToField(tt.text)

			assert.NoError(t, err, "エラーが発生していないこと")
			assert.Equal(t, tt.want, result)
		})
	}
}
