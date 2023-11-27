package article

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	UniqueId     UniqueId
	UserUniqueId string
	Title        Title
	Content      Content
	UpdatedAt    time.Time
	CreatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func NewArticle(
	UniqueId UniqueId,
	UserUniqueId string,
	Title Title,
	Content Content) *Article {
	newArticle := Article{
		UniqueId:     UniqueId,
		UserUniqueId: UserUniqueId,
		Title:        Title,
		Content:      Content,
	}
	return &newArticle
}

func CreateArticle(
	userUniqueId string,
	title string,
	content string,
) (*Article, error) {
	newUniqueId, err := NewUniqueId()
	if err != nil {
		return nil, err
	}

	newTitle, err := NewTitle(title)
	if err != nil {
		return nil, err
	}
	newContent, err := NewContent(content)
	if err != nil {
		return nil, err
	}

	newArticle := NewArticle(
		*newUniqueId,
		userUniqueId,
		*newTitle,
		*newContent,
	)

	return newArticle, nil
}

func (article *Article) UpdateTitleAndContent(
	title string,
	content string,
) error {
	newTitle, err := NewTitle(title)
	if err != nil {
		return err
	}
	newContent, err := NewContent(content)
	if err != nil {
		return err
	}
	article.Content = *newContent
	article.Title = *newTitle

	return nil
}
