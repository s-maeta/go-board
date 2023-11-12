package usecase

import (
	preUserModel "board/app/domain/model/pre_user"
	"board/app/infrastructure/repository"
	"board/app/infrastructure/service"
	request "board/app/interface/request/user"
	"board/config"
	"fmt"
)

type UserRegisterUseCase struct {
	PreUserRepository preUserModel.IPreUserRepository
}

func NewUserRegisterUseCase() *UserRegisterUseCase {
	return &UserRegisterUseCase{
		PreUserRepository: repository.NewPreUserRepository(),
	}
}

func (useCase *UserRegisterUseCase) Execute(request *request.UserRegisterRequest) error {

	//仮登録ユーザーエンティティ生成
	preUser, err := preUserModel.CreatePreUser(
		request.Email,
		request.Password,
	)
	if err != nil {
		return err
	}

	// 仮ユーザーDB永続化
	err = useCase.PreUserRepository.Create(*preUser)
	if err != nil {
		return err
	}

	// 仮ユーザー登録完了メール送信
	config := config.GetConfig()

	err = service.SendEmail(
		"test@test.go",
		[]string{string(preUser.Email)},
		"ユーザー仮登録完了",
		fmt.Sprintf("ユーザー仮登録が完了しました。以下のメールアドレスをクリックして本登録を完了してください。/n%v/user/create/%v",
			config.ApiUrl,
			preUser.Token),
	)
	if err != nil {
		return err
	}

	return nil
}
