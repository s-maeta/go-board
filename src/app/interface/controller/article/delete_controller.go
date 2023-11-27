package article

import (
	"board/app/interface/controller"
	request "board/app/interface/request/article"
	useCase "board/app/usecase/article"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteController struct {
	controller controller.Controller
	useCase    useCase.DeleteUseCase
}

func NewDeleteController() *DeleteController {
	return &DeleteController{
		controller: controller.Controller{},
		useCase:    *useCase.NewDeleteUseCase(),
	}
}

func (controller *DeleteController) Handler(context *gin.Context) {
	request := request.DeleteRequest{
		ArticleId: context.Param("id"),
	}
	err := controller.useCase.Execute(&request)
	if err != nil {
		controller.controller.ErrorJson(context, http.StatusBadRequest, []string{err.Error()})
	}

	context.Status(http.StatusCreated)
}
