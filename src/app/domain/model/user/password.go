package user

import "board/app/infrastructure/service"

type Password string

func NewPassword(password string) (*Password, error) {
	//セキュリティに関する暗号化複合化はドメイン知識とするべきなのか検討中、、、、
	newPassword := Password(password)
	service.Encrypt(string(newPassword))
	return &newPassword, nil
}
