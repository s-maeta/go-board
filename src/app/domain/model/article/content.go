package article

import "errors"

type Content string

func NewContent(content string) (*Content, error) {
	if len(content) >= 2000 {
		return nil, errors.New("コンテンツは2000文字以内で入力してください。")
	}
	newContent := Content(content)

	return &newContent, nil
}
