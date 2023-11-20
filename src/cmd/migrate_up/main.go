package main

import (
	"board/config"
	"board/database"
	"fmt"
	"log"
)

func main() {
	config.Init()

	database.Init()
	defer database.Close()

	m := database.GetM()

	// マイグレーションの実行
	err := m.Up()
	// エラーメッセージが「no change」の場合もスキップ
	if err != nil && err.Error() != "no change" {
		log.Printf("m.Up() Error Message: %s\n", err)
		fmt.Println("migrate up failed")
	}
}
