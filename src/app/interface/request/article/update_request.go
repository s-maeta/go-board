package article

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type UpdateRequest struct {
	ArticleId string
	Title     string `json:title`
	Content   string `json:content`
}

func (request *UpdateRequest) Validate() {
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
	// err := validation.ValidateStruct(&request, validation.Field(&request.ArticleId,
	// 	validation.Required.Error("記事IDは必ず指定してください。"),
	// ),
	// )
	// if err != nil {
	// 	return err
	// }
	// return nil
}
