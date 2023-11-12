package repository

import (
	"board/app/domain/model/pre_user"
	"board/app/infrastructure/dto"
	"board/database"

	"gorm.io/gorm"
)

type PreUserRepository struct {
	db gorm.DB
}

func NewPreUserRepository() pre_user.IPreUserRepository {
	db := database.GetDB()
	return &PreUserRepository{
		db: *db,
	}
}

func (repository *PreUserRepository) Create(preUser pre_user.PreUser) error {
	result := repository.db.Create(preUser)
	if result != nil {
		return result.Error
	}
	return nil
}

func (repository *PreUserRepository) FindForToken(token pre_user.Token) (pre_user.PreUser, error) {
	preUserDto := dto.PreUserDto{}
	repository.db.Where("token = ?", token).First(&preUserDto)

	preUser := preUserDto.ToEntity()
	return *preUser, nil
}
