package repository

import (
	"board/app/domain/model/article"
	"board/database"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository() *ArticleRepository {
	db := database.GetDB()
	return &ArticleRepository{
		db: db,
	}
}

func (repository *ArticleRepository) Create(article *article.Article) error {
	result := repository.db.Create(article)
	if result != nil {
		return result.Error
	}
	return nil
}
