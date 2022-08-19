package resparse

import (
	"contentful-article-cli/scripts/response"
	"encoding/json"
	"fmt"
)

// FieldToString はContentfulのFieldを文字列に変換する
func FieldToString(field interface{}) (string, error) {
	byte, err := json.Marshal(field)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var body response.Lang
	if err := json.Unmarshal(byte, &body); err != nil {
		fmt.Println(err)
		return "", err
	}

	return body.Ja, nil
}

// StringToField は文字列をContentfulのFieldに変換する
func StringToField(text string) (interface{}, error) {
	byte, err := json.Marshal(&response.Lang{Ja: text})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var body map[string]interface{}
	if err := json.Unmarshal(byte, &body); err != nil {
		fmt.Println(err)
		return "", err
	}

	return body, nil
}
