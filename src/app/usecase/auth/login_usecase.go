package auth

import (
	"board/app/domain/model/user"
	"board/app/infrastructure/repository"
	"board/app/infrastructure/service"
	authRequest "board/app/interface/request/auth"
	authResponse "board/app/interface/response/auth"
	"errors"
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

	user := useCase.UserRepository.FindForEmail(
		user.Email(params.Email),
	)

	//パスワード一致チェック
	isVerifyPassword := service.VerifyPassword(params.Password, string(user.Password))
	if !isVerifyPassword {
		return nil, errors.New("パスワード認証に失敗しました。")
	}

	// トークン生成
	token, err := service.GenerateToken(string(user.UniqueId))
	if err != nil {
		return nil, err
	}

	return &authResponse.LoginResponse{Token: token}, nil
}
