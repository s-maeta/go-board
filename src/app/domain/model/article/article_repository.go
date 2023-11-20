package article

type IArticleRepository interface {
	Create(article *Article) error
}
