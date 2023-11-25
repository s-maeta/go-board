package dto

import "board/app/domain/model/article"

type ArticleDto struct {
	UniqueId     string `gorm:"primaryKey"`
	UserUniqueId string
	Title        string
	Content      string
}

// テーブル名を指定
func (ArticleDto ArticleDto) TableName() string {
	return "articles"
}

// DTOからEntityに変換する
func (articleDto *ArticleDto) ToEntity() *article.Article {
	entity := article.NewArticle(
		article.UniqueId(articleDto.UniqueId),
		articleDto.UserUniqueId,
		article.Title(articleDto.Title),
		article.Content(articleDto.Content),
	)
	return entity
}
