package article

import "board/app/domain/model/article"

type ShowResponse struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (response *ShowResponse) ToResponse(article *article.Article) {
	response.Content = string(article.Content)
	response.Title = string(article.Title)
}
