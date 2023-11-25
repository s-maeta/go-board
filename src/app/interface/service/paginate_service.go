package service

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type PaginateService struct {
	Page   int `from:page json:page`
	Number int `from:number json:number`
}

// PaginateQuery ページネーション用のクエリを返す
func (service *PaginateService) PaginateQuery(db *gorm.DB) *gorm.DB {
	return db.Limit(service.Number).Offset(service.Page)
}

func (service *PaginateService) Validate() error {
	err := validation.ValidateStruct(service,
		validation.Field(&service.Number, validation.Required.Error("取得件数は必須です。")),
		validation.Field(&service.Page, validation.Required.Error("page数は必須です。")),
	)
	if err != nil {
		return err
	}
	return nil
}

// func (service *PaginateService) ValidationRules() []*validation.FieldRules {
// 	var rules []*validation.FieldRules

// 	// Numberに関するルール
// 	numberRules :=validation.Field(service.Number, validation.Required.Error("取得件数は必須です。"))
// 	rules = append(rules, numberRules)

// 	// Pageに関するルール
// 	pageRules := validation.Field(service.Page, validation.Required.Error("page数は必須です。"))
// 	rules = append(rules, pageRules)

// 	return rules
// }
