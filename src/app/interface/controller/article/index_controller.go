package article

import (
	"board/app/interface/controller"
	request "board/app/interface/request/article"
	"board/app/interface/service"
	useCase "board/app/usecase/article"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
	controller *controller.Controller
	useCase    *useCase.IndexUseCase
}

func NewIndexController() *IndexController {
	return &IndexController{
		controller: &controller.Controller{},
		useCase:    useCase.NewIndexUseCase(),
	}
}

func (controller *IndexController) Handler(context *gin.Context) {
	page, err := strconv.Atoi(context.Query("page"))
	number, err := strconv.Atoi(context.Query("number"))

	if err != nil {
		controller.controller.ErrorJson(context, http.StatusUnprocessableEntity, []string{err.Error()})
		return
	}

	request := request.IndexRequest{
		PaginateService: &service.PaginateService{
			Page:   page,
			Number: number,
		},
	}
	// err := context.ShouldBindQuery(&request)
	if err = request.Validate(); err != nil {
		controller.controller.ErrorJson(context, http.StatusUnprocessableEntity, []string{err.Error()})
		return
	}

	response, err := controller.useCase.Execute(request)

	if err != nil {
		controller.controller.ErrorJson(context, http.StatusBadRequest, []string{"error"})
		return
	}
	context.JSON(http.StatusOK, response)
}
