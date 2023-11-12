package user

import (
	"board/app/interface/controller"
	"board/app/interface/request/user"
	"board/app/usecase"

	"github.com/gin-gonic/gin"
)

type UserCreateController struct {
	Controller *controller.Controller
	UseCase    *usecase.UserCreateUseCase
}

func NewUserCreateController() *UserCreateController {
	return &UserCreateController{
		Controller: &controller.Controller{},
		UseCase:    usecase.NewUserCreateUseCase(),
	}
}

func (userController *UserCreateController) Handler(context *gin.Context) {
	// var request user.UserCreateRequest
	token := context.Param("token")

	request := user.UserCreateRequest{Token: token}
	//bindJsonでなければバリデーションが自動実行されないため、Validateを呼び出し
	request.Validate()

	err := userController.UseCase.Execute(&request)
	if err != nil {
		userController.Controller.ErrorJson(context, 500, []string{err.Error()})
	}
}
