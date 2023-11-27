package article

import (
	"board/app/interface/controller"
	request "board/app/interface/request/article"
	useCase "board/app/usecase/article"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShowController struct {
	controller controller.Controller
	useCase    useCase.ShowUseCase
}

func NewShowController() *ShowController {
	return &ShowController{
		controller: controller.Controller{},
		useCase:    *useCase.NewShowUseCase(),
	}
}

func (controller *ShowController) Handler(context *gin.Context) {
	request := request.ShowRequest{
		ArticleId: context.Param("id"),
	}

	response, err := controller.useCase.Execute(&request)
	if err != nil {
		controller.controller.ErrorJson(context, http.StatusBadRequest, []string{err.Error()})
		return
	}
	context.JSON(http.StatusOK, response)
}
