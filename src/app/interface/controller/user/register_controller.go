package user

import (
	"board/app/interface/controller"
	"board/app/interface/request/user"
	"board/app/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRegisterController struct {
	Controller *controller.Controller
	UseCase    *usecase.UserRegisterUseCase
}

func NewUserRegisterController() *UserRegisterController {
	return &UserRegisterController{
		Controller: &controller.Controller{},
		UseCase:    usecase.NewUserRegisterUseCase(),
	}
}

func (controller *UserRegisterController) Handler(context *gin.Context) {
	request := &user.UserRegisterRequest{}

	err := context.ShouldBindJSON(&request)
	if err != nil {
		controller.Controller.ErrorJson(context, 400, []string{"Parameters are not found.", err.Error()})
	}

	err = controller.UseCase.Execute(request)
	if err != nil {
		controller.Controller.ErrorJson(context, 400, []string{err.Error()})
	}

	context.Status(http.StatusOK)
}
