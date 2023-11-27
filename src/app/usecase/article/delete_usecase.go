package article

import (
	articleModel "board/app/domain/model/article"
	"board/app/infrastructure/repository"
	request "board/app/interface/request/article"
	"errors"
)

type DeleteUseCase struct {
	articleRepository articleModel.IArticleRepository
}

func NewDeleteUseCase() *DeleteUseCase {
	return &DeleteUseCase{
		articleRepository: repository.NewArticleRepository(),
	}
}

func (useCase *DeleteUseCase) Execute(params *request.DeleteRequest) error {

	article, err := useCase.articleRepository.Find(articleModel.UniqueId(params.ArticleId))
	if err != nil {
		return err
	}
	if article == nil {
		return errors.New("削除対象のレコードが見つかりませんでした。")
	}

	err = useCase.articleRepository.Delete(article)
	if err != nil {
		return err
	}

	return nil
}
