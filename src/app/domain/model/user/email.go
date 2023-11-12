package user

import (
	"errors"
	"regexp"
)

type Email string

func NewEmail(email string) (*Email, error) {
	if err := validate(email); err != nil {
		return nil, err
	}

	newEmail := Email(email)
	return &newEmail, nil
}

func validate(email string) error {
	// 正規表現チェック
	regexPattern := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !regexPattern.MatchString(email) {
		return errors.New("メールアドレスの形式が間違っています。")
	}

	return nil
}
