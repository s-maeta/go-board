package auth

import validation "github.com/go-ozzo/ozzo-validation"

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:password binding:"required"`
}

func (request *LoginRequest) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(
			&request.Password,
			validation.Required.Error("パスワードは必ず入力してください。"),
		),
		validation.Field(
			&request.Email,
			validation.Required.Error("メールアドレスは必ず入力してください。"),
		),
	)
}
