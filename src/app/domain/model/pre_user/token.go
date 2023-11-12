package pre_user

import "board/app/infrastructure/service"

type Token string

func NewToken(UniqueId UniqueId) (*Token, error) {
	token, err := service.GenerateToken(string(UniqueId))
	if err != nil {
		return nil, err
	}
	newToken := Token(token)
	return &newToken, nil
}
