package article

import (
	"board/app/domain/model/article"
	"board/app/infrastructure/repository"
	request "board/app/interface/request/article"
	response "board/app/interface/response/article"
)

type IndexUseCase struct {
	articleRepository article.IArticleRepository
}

func NewIndexUseCase() *IndexUseCase {
	return &IndexUseCase{
		articleRepository: repository.NewArticleRepository(),
	}
}

func (useCase *IndexUseCase) Execute(request request.IndexRequest) (*response.IndexResponse, error) {
	articles, err := useCase.articleRepository.Index(*request.PaginateService)
	if err != nil {
		return nil, err
	}

	var response response.IndexResponse
	response.ToResponse(articles)

	return &response, nil

}
