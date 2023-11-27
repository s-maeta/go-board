package article

import (
	articleModel "board/app/domain/model/article"
	"board/app/infrastructure/repository"
	request "board/app/interface/request/article"
	"errors"
)

type UpdateUseCase struct {
	articleRepository articleModel.IArticleRepository
}

func NewUpdateUseCase() *UpdateUseCase {
	return &UpdateUseCase{
		articleRepository: repository.NewArticleRepository(),
	}
}

func (useCase *UpdateUseCase) Execute(params *request.UpdateRequest) error {

	article, err := useCase.articleRepository.Find(articleModel.UniqueId(params.ArticleId))
	if err != nil {
		return err
	}
	if article == nil {
		return errors.New("更新対象のレコードが見つかりませんでした。")
	}

	err = article.UpdateTitleAndContent(params.Title, params.Content)
	if err != nil {
		return err
	}
	err = useCase.articleRepository.Update(article)
	if err != nil {
		return err
	}

	return nil
}
