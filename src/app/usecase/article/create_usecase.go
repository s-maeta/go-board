package article

import (
	"board/app/domain/model/article"
	"board/app/domain/model/user"
	"board/app/infrastructure/repository"
	request "board/app/interface/request/article"
)

type CreateUseCase struct {
	articleRepository article.IArticleRepository
}

func NewCreateUseCase() *CreateUseCase {
	return &CreateUseCase{
		articleRepository: repository.NewArticleRepository(),
	}
}

func (useCase *CreateUseCase) Execute(request *request.CreateRequest, authUser *user.User) error {
	//記事を作成する
	createArticle, err := article.CreateArticle(
		string(authUser.UniqueId),
		request.Title,
		request.Content,
	)
	if err != nil {
		return err
	}
	err = useCase.articleRepository.Create(createArticle)
	if err != nil {
		return err
	}
	return nil
}
