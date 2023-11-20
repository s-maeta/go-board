package article

import validation "github.com/go-ozzo/ozzo-validation"

type CreateRequest struct {
	Title   string `json:title`
	Content string `json:content`
}

func (request *CreateRequest) Validate() {
	validation.ValidateStruct(&request,
		validation.Field(&request.Title,
			validation.Required.Error("タイトルは必須で入力してください。"),
			validation.Length(1, 50),
		),
		validation.Field(&request.Content,
			validation.Required.Error("投稿内容は必ず入力してください。"),
			validation.Length(1, 1000),
		),
	)
}
