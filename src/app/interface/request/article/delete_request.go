package article

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type DeleteRequest struct {
	ArticleId string
}

func (request *DeleteRequest) Validate() error {
	err := validation.ValidateStruct(&request, validation.Field(&request.ArticleId,
		validation.Required.Error("記事IDは必ず指定してください。"),
	),
	)
	if err != nil {
		return err
	}
	return nil
}
