package repository

import (
	userModel "board/app/domain/model/user"
	"board/app/infrastructure/dto"
	"board/database"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() userModel.IUserRepository {
	db := database.GetDB()
	return &UserRepository{
		db: db,
	}
}

func (repository *UserRepository) Create(user *userModel.User) error {
	result := repository.db.Create(user)
	if result != nil {
		return result.Error
	}
	return nil
}

func (repository *UserRepository) Delete(user *userModel.User) error {
	var userDto dto.UserDto
	result := repository.db.Where("unique_id = ?", user.UniqueId).Delete(&userDto)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repository *UserRepository) FindForLoginUser(email userModel.Email, password userModel.Password) *userModel.User {
	var userDto dto.UserDto
	result := repository.db.Where("password = ?", password).Where("email = ?", email).First(&userDto)

	if result.Error == gorm.ErrRecordNotFound {
		return nil
	}
	// if result.Error != nil {
	// 	panic("SQL Error")
	// }

	user := userDto.ToEntity()

	return &user
}

func (repository *UserRepository) FindForUniqueId(uniqueId string) *userModel.User {
	var userDto dto.UserDto

	result := repository.db.Where("unique_id = ?", uniqueId).First(*&userDto)

	if result.Error != nil {
		return nil
	}
	user := userDto.ToEntity()
	return &user

}
