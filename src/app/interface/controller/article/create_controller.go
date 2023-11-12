package article

import (
	"board/app/interface/controller"
	request "board/app/interface/request/article"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateController struct{
	controller *controller.Controller
}

func NewCreateController ()*CreateController{
	return &CreateController{
		controller: &controller.Controller{},
	}
}

func (controller *CreateController) Handler(context *gin.Context){
	request := request.CreateRequest{}

	err := context.ShouldBindJSON(&request)

	if err != nil {
		controller.controller.ErrorJson(context,http.StatusBadRequest,[]string{err.Error()})
	}


}