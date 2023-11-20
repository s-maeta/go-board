package user

import (
	"board/app/domain/model/pre_user"
	"errors"
	"time"

	"gorm.io/gorm"
)

type User struct {
	UniqueId  UniqueId
	Name      Name
	Password  Password
	Email     Email
	UpdatedAt time.Time
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// コンストラクタ
func NewUser(
	unique_user_id UniqueId,
	name Name,
	password Password,
	email Email,
) *User {
	newUser := User{
		UniqueId: unique_user_id,
		Name:     name,
		Password: password,
		Email:    email,
	}
	return &newUser
}

func CreateUser(
	PreUser *pre_user.PreUser,
) (*User, error) {
	newUniqueId, err := NewUniqueId()
	if err != nil {
		return nil, err
	}
	newName, err := NewName("test")
	if err != nil {
		return nil, err
	}
	newPassword := Password(string(PreUser.Password))
	if err != nil {
		return nil, err
	}
	newEmail, err := NewEmail(string(PreUser.Email))
	if err != nil {
		return nil, err
	}

	user := User{
		UniqueId: *newUniqueId,
		Name:     *newName,
		Password: newPassword,
		Email:    *newEmail,
	}
	return &user, nil
}

func (User *User) Exists(subjectUser User) error {
	if subjectUser.Email == User.Email && subjectUser.Password == User.Password {
		return errors.New("user is exists")
	}
	return nil
}
