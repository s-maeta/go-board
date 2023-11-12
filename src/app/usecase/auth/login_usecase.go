package auth

import (
	"board/app/domain/model/user"
	"board/app/infrastructure/repository"
	"board/app/infrastructure/service"
	authRequest "board/app/interface/request/auth"
	authResponse "board/app/interface/response/auth"
)

type LoginUseCase struct {
	UserRepository user.IUserRepository
}

func NewLoginUseCase() *LoginUseCase {
	return &LoginUseCase{
		UserRepository: repository.NewUserRepository(),
	}
}

func (useCase *LoginUseCase) Execute(params *authRequest.LoginRequest) (*authResponse.LoginResponse, error) {

	// ユーザー存在チェック
	user := useCase.UserRepository.FindForLoginUser(
		user.Email(params.Email),
		user.Password(params.Password),
	)

	//仮登録ユーザーを削除する。

	token, err := service.GenerateToken(string(user.UniqueId))
	if err != nil {
		return nil, err
	}

	return &authResponse.LoginResponse{Token: token}, nil
}
