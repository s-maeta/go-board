package router

import (
	"board/app/infrastructure/middleware"
	"board/app/infrastructure/service"
	"board/app/interface/controller/article"
	"board/app/interface/controller/auth"
	"board/app/interface/controller/user"
	"net/http"

	"github.com/gin-contrib/cors"
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
	articleIndexController := article.NewIndexController()
	articleShowController := article.NewShowController()
	articleDeleteController := article.NewDeleteController()
	articleUpdateController := article.NewUpdateController()

	//GINのデフォルトバリデーション（go-playground）を独自バリデーションに上書き
	binding.Validator = service.NewOzzoValidator()

	//CORSの設定
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		// MaxAge: 24 * time.Hour,
	}))

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

	router.GET("user/", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"message": "test"}) })

	article := router.Group("/article")
	{
		article.POST("/create", articleCreateController.Handler)
		article.GET("/index", articleIndexController.Handler)
		article.GET("/:id", articleShowController.Handler)
		article.DELETE("/:id", articleDeleteController.Handler)
		article.PUT(":id", articleUpdateController.Handler)
	}

	router.Run(":3001")
}
