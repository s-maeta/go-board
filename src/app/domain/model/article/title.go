package article

import "errors"

type Title string

func NewTitle(title string) (*Title, error) {
	if len(title) >= 50 {
		return nil, errors.New("タイトルは50文字以内で入力してください。")
	}
	newTitle := Title(title)
	return &newTitle, nil
}
