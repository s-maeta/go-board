package article

import "board/app/interface/service"

type IArticleRepository interface {
	Create(article *Article) error
	Index(service.PaginateService) ([]Article, error)
}
