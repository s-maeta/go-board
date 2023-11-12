package dto

import "board/app/domain/model/user"

type UserDto struct {
	UniqueId string `gorm:"primaryKey"`
	Name     string
	Password string
	Email    string
}

func (UserDto UserDto) TableName() string {
	return "users"
}

func (UserDto UserDto) ToEntity() user.User {
	return *user.NewUser(
		user.UniqueId(UserDto.UniqueId),
		user.Name(UserDto.Name),
		user.Password(UserDto.Password),
		user.Email(UserDto.Email),
	)
}
