package user

import validation "github.com/go-ozzo/ozzo-validation"

type UserCreateRequest struct {
	Token string `json:token`
}

func (request *UserCreateRequest) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(
			&request.Token,
			validation.Required.Error("トークンは必ず入力してください"),
		),
	)
}
