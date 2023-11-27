package article

import (
	"board/app/interface/controller"
	request "board/app/interface/request/article"
	useCase "board/app/usecase/article"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateController struct {
	controller controller.Controller
	useCase    useCase.UpdateUseCase
}

func NewUpdateController() *UpdateController {
	return &UpdateController{
		controller: controller.Controller{},
		useCase:    *useCase.NewUpdateUseCase(),
	}
}

func (controller *UpdateController) Handler(context *gin.Context) {
	var request request.UpdateRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		controller.controller.ErrorJson(context, http.StatusBadRequest, []string{err.Error()})
	}
	request.ArticleId = context.Param("id")

	err := controller.useCase.Execute(&request)
	if err != nil {
		controller.controller.ErrorJson(context, http.StatusBadRequest, []string{err.Error()})
	}

	context.Status(http.StatusCreated)

}
