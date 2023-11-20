package pre_user

import (
	"time"

	"gorm.io/gorm"
)

type PreUser struct {
	UniqueId  UniqueId
	Email     Email
	Password  Password
	Token     Token
	UpdatedAt time.Time
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func NewPreUser(
	UniqueId UniqueId,
	Email Email,
	Password Password,
	Token Token,
) *PreUser {
	return &PreUser{
		UniqueId: UniqueId,
		Email:    Email,
		Password: Password,
		Token:    Token,
	}
}

func CreatePreUser(
	Email,
	Password string,
) (*PreUser, error) {

	uniqueId, err := NewUniqueId()
	if err != nil {
		return nil, err
	}

	email, err := NewEmail(Email)
	if err != nil {
		return nil, err
	}

	password, err := NewPassword(Password)
	if err != nil {
		return nil, err
	}

	token, err := NewToken(*uniqueId)
	if err != nil {
		return nil, err
	}

	newPreUser := PreUser{
		UniqueId: *uniqueId,
		Email:    *email,
		Password: *password,
		Token:    *token,
	}

	return &newPreUser, nil
}
