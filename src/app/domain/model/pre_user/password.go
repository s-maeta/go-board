package pre_user

import "board/app/infrastructure/service"

type Password string

func NewPassword(password string) (*Password, error) {
	//セキュリティに関する暗号化複合化はドメイン知識とするべきなのか検討中、、、、
	encryptPass, err := service.Encrypt(string(password))
	if err != nil {
		return nil, err
	}
	newPassword := Password(encryptPass)
	return &newPassword, nil
}
