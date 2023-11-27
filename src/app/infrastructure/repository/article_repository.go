package repository

import (
	"board/app/domain/model/article"
	"board/app/infrastructure/dto"
	"board/app/interface/service"
	"board/database"
	"errors"

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

func (repository *ArticleRepository) Find(articleId article.UniqueId) (*article.Article, error) {
	articleDto := dto.ArticleDto{}

	result := repository.db.Where("unique_id = ?", articleId).First(&articleDto)
	//レコードが見つからなかった場合にはnilを返却する
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}

	article := articleDto.ToEntity()
	return article, nil
}

func (repository *ArticleRepository) Delete(article *article.Article) error {
	var articleDto dto.ArticleDto
	result := repository.db.Where("unique_id = ?", article.UniqueId).Delete(&articleDto)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
