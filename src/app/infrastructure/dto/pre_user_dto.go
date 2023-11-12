package dto

import "board/app/domain/model/pre_user"

type PreUserDto struct {
	UniqueId string `gorm:"primaryKey"`
	Email    string
	Password string
	Token    string
}

// テーブル名を指定
func (preUserDto PreUserDto) TableName() string {
	return "pre_users"
}

// DTOからEntityに変換する
func (preUserDto PreUserDto) ToEntity() *pre_user.PreUser {
	entity := pre_user.NewPreUser(
		pre_user.UniqueId(preUserDto.UniqueId),
		pre_user.Email(preUserDto.Email),
		pre_user.Password(preUserDto.Password),
		pre_user.Token(preUserDto.Token),
	)
	return entity
}
