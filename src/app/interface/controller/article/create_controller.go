package article

import (
	"board/app/interface/controller"
	request "board/app/interface/request/article"
	use_case "board/app/usecase/article"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateController struct {
	controller *controller.Controller
	useCase    *use_case.CreateUseCase
}

func NewCreateController() *CreateController {
	return &CreateController{
		controller: &controller.Controller{},
		useCase:    use_case.NewCreateUseCase(),
	}
}

func (controller *CreateController) Handler(context *gin.Context) {
	request := request.CreateRequest{}

	err := context.ShouldBindJSON(&request)

	if err != nil {
		controller.controller.ErrorJson(context, http.StatusBadRequest, []string{err.Error()})
	}

	authUser := controller.controller.AuthUser(context)

	err = controller.useCase.Execute(&request, authUser)
	if err != nil {
		controller.controller.ErrorJson(context, 500, []string{err.Error()})
		return
	}
}
