package router

import (
	"board/app/infrastructure/middleware"
	"board/app/infrastructure/service"
	"board/app/interface/controller/article"
	"board/app/interface/controller/auth"
	"board/app/interface/controller/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Init() {
	router := gin.Default()

	//コントローラーの作成
	userCreateController := user.NewUserCreateController()
	userRegisterController := *user.NewUserRegisterController()
	authController := auth.NewLoginController()
articleCreateController := article.NewCreateController()
	//GINのデフォルトバリデーション（go-playground）を独自バリデーションに上書き
	binding.Validator = service.NewOzzoValidator()

	//ルーティングがなかった場合のルート
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"errors": "route not found"})
	})

	router.POST("/login", authController.Login)

	user := router.Group("/user")
	{
		user.POST("/register", userRegisterController.Handler)
		user.POST("/create/:token", userCreateController.Handler)
	}

	router.Use(middleware.LoginCheckMiddleware)

	router.GET("user/", func(ctx *gin.Context) { ctx.JSON(200,gin.H{"message":"test",})})

	article :=router.Group("/article")
	{
		article.POST("create",articleCreateController.Handler)
	}

	router.Run(":3001")
}
