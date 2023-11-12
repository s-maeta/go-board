package auth

import (
	"board/app/interface/controller"
	authRequest "board/app/interface/request/auth"
	loginUseCase "board/app/usecase/auth"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	Controller *controller.Controller
}

func NewLoginController() *LoginController {
	return &LoginController{
		Controller: &controller.Controller{},
	}
}

func (controller *LoginController) Login(context *gin.Context) {
	var request *authRequest.LoginRequest
	err := context.ShouldBindJSON(&request)

	if err != nil {
		controller.Controller.ErrorJson(context, 500, []string{err.Error()})
	}

	result, err := loginUseCase.NewLoginUseCase().Execute(request)

	if err != nil {
		controller.Controller.ErrorJson(context, 500, []string{err.Error()})
	}
	context.JSON(200, result)
}
