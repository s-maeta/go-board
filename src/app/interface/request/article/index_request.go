package article

import (
	"board/app/interface/service"
)

type IndexRequest struct {
	*service.PaginateService
}

func (request *IndexRequest) Validate() error {
	err := request.PaginateService.Validate()
	if err != nil {
		return err
	}
	return nil
}
