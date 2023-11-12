package user

import validation "github.com/go-ozzo/ozzo-validation"

type UserRegisterRequest struct {
	Password string `json:password`
	Email    string `json:email`
}

func (request *UserRegisterRequest) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(
			&request.Password,
			validation.Required.Error("パスワードは必ず入力してください。"),
		),
		validation.Field(
			&request.Password,
			validation.Required.Error("メールアドレスは必ず入力してください。"),
		),
	)
}
