package repository

import (
	"board/app/domain/model/article"
	"board/app/interface/service"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestArticleRepository_Index(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		service service.PaginateService
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []article.Article
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &ArticleRepository{
				db: tt.fields.db,
			}
			got, err := repository.Index(tt.args.service)
			if (err != nil) != tt.wantErr {
				t.Errorf("ArticleRepository.Index() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArticleRepository.Index() = %v, want %v", got, tt.want)
			}
		})
	}
}
