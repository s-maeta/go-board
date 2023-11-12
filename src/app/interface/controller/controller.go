package controller

import (
	"github.com/gin-gonic/gin"
)

// コントローラーの規定Struct
type Controller struct{}

// type JsonResponse interface{
// 	ToJson()()
// }

// エラーを返却する
func (c *Controller) ErrorJson(context *gin.Context, status int, errors []string) {
	context.JSON(status, gin.H{"errors": errors})
}

// // 正常レスポンスを返却する
// func(c *Controller) SuccessJson(context *gin.Context,status int,data JsonResponse ){
// 	context.JSON(status,gin.H(data.ToJson()))
// }
