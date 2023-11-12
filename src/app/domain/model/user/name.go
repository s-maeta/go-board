package user

import "errors"

type Name string

func NewName(name string) (*Name, error) {
	if len(name) > 255 {
		return nil, errors.New("指名は255文字以内で指定してください。")
	}
	newName := Name(name)
	return &newName, nil
}
