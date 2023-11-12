package user

import (
	"board/app/interface/controller"
	"board/app/interface/request/user"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Controller *controller.Controller
}

func NewUserController() *UserController {
	return &UserController{
		Controller: &controller.Controller{},
	}
}

func (userController *UserController) GetUser(context *gin.Context) {
	userId, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		userController.Controller.ErrorJson(context, 400, []string{"Parameters are not found."})
	}

	// リクエストを作成
	request := &user.UserGetRequest{UserId: userId}
	fmt.Println(request)

	// useCaseで処理を実行

	if err != nil {
		userController.Controller.ErrorJson(context, 400, []string{err.Error()})
		return
	}

	context.JSON(200, "これはテストです。")
}
