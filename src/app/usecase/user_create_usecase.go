package usecase

import (
	preUserModel "board/app/domain/model/pre_user"
	userModel "board/app/domain/model/user"
	"board/app/infrastructure/repository"
	request "board/app/interface/request/user"
)

type UserCreateUseCase struct {
	UserRepository    userModel.IUserRepository
	PreUserRepository preUserModel.IPreUserRepository
}

func NewUserCreateUseCase() *UserCreateUseCase {
	return &UserCreateUseCase{
		UserRepository:    repository.NewUserRepository(),
		PreUserRepository: repository.NewPreUserRepository(),
	}
}

func (useCase *UserCreateUseCase) Execute(request *request.UserCreateRequest) error {

	preUser, err := useCase.PreUserRepository.FindForToken(preUserModel.Token(request.Token))
	if err != nil {
		return err
	}

	// 新規ユーザー作成
	user, err := userModel.CreateUser(&preUser)
	if err != nil {
		return err
	}

	// ユーザー存在確認
	currentUser := useCase.UserRepository.FindForLoginUser(user.Email, user.Password)
	err = user.Exists(*currentUser)
	if err != nil {
		return err
	}

	useCase.UserRepository.Create(user)

	return nil
}
