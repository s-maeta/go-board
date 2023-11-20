package controller

import (
	"board/app/domain/model/user"
	"net/http"

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
	// エラーが発生した場合、即座に処理を中断する
	context.Abort()
}

func (c *Controller) AuthUser(context *gin.Context) *user.User {
	authUser, ok := context.Get("user")
	if !ok {
		c.ErrorJson(context, http.StatusBadRequest, []string{"認証ユーザーの取得に失敗しました。"})
		// エラーが発生した場合、即座に処理を中断する
		context.Abort()
	}
	//ポインタ型でアサーションする。
	u, ok := authUser.(*user.User)

	if !ok {
		c.ErrorJson(context, http.StatusBadRequest, []string{"Userの型アサーションに失敗しました。"})
		// エラーが発生した場合、即座に処理を中断する
		context.Abort()
	}

	return u
}

// // 正常レスポンスを返却する
// func(c *Controller) SuccessJson(context *gin.Context,status int,data JsonResponse ){
// 	context.JSON(status,gin.H(data.ToJson()))
// }
