package article

import (
	articleModel "board/app/domain/model/article"
	"board/app/infrastructure/repository"
	request "board/app/interface/request/article"
	response "board/app/interface/response/article"
	"errors"
)

type ShowUseCase struct {
	articleRepository articleModel.IArticleRepository
}

func NewShowUseCase() *ShowUseCase {
	return &ShowUseCase{
		articleRepository: repository.NewArticleRepository(),
	}
}

func (useCase *ShowUseCase) Execute(params *request.ShowRequest) (*response.ShowResponse, error) {
	articleId := articleModel.UniqueId(params.ArticleId)
	article, err := useCase.articleRepository.Find(articleId)
	if err != nil {
		return nil, err
	}
	if article == nil {
		return nil, errors.New("検索対象のレコードが見つかりませんでした。")
	}

	var res response.ShowResponse
	res.ToResponse(article)

	return &res, nil
}
