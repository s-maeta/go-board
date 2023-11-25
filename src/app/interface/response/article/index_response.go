package article

import "board/app/domain/model/article"

type IndexResponse struct {
	Items []Item `json:"items"`
}
type Item struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (response *IndexResponse) ToResponse(articles []article.Article) {
	for _, article := range articles {
		response.Items = append(response.Items, Item{
			Title:   string(article.Title),
			Content: string(article.Content),
		})
	}
}
