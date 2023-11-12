package article

import validation "github.com/go-ozzo/ozzo-validation"

type CreateRequest struct {
	title string
	content string
}

func (request *CreateRequest) Validate(){
	validation.ValidateStruct(&request,
	validation.Field(
		&request.title,
		validation.Required.Error("タイトルは必須で入力してください。"),
		validation.Length(1,50),
	),
	validation.Field(&request.content,
		validation.Required.Error("投稿内容は必ず入力してください。"),
		validation.Length(1,1000),
	),
)
}
