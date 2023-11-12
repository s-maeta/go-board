package main

import (
	"board/config"
	"board/database"
	"board/router"
)

func main() {
	// config の初期化
	config.Init()

	// DBの初期化
	database.Init()
	defer database.Close()

	// //GINのデフォルトバリデーション（go-playground）を独自バリデーションに上書き
	// binding.Validator = service.NewOzzoValidator()

	//ルーティングの初期化
	router.Init()
}
