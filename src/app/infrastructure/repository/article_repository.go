package repository

import (
	"board/app/domain/model/article"
	"board/app/infrastructure/dto"
	"board/app/interface/service"
	"board/database"
	"fmt"

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

func (repository *ArticleRepository) Index(service service.PaginateService) ([]article.Article, error) {
	fmt.Println(service.Number)
	articleDto := []dto.ArticleDto{}
	query := service.PaginateQuery(repository.db.Find(&articleDto))
	query.Debug().Find(&articleDto)

	var articles []article.Article
	for _, value := range articleDto {
		articles = append(articles, *value.ToEntity())
	}
	return articles, nil
}

// func (repository *ArticleRepository) Index(service service.PaginateService) ([]article.Article, error) {
// 	articleDto := []dto.ArticleDto{}
// 	repository.db.Debug().Find(&articleDto)

// 	var articles []article.Article
// 	for _, value := range articleDto {
// 		articles = append(articles, *value.ToEntity())
// 	}
// 	return articles, nil
// }
