package main

import (
	"board/config"
	"board/database"
	"fmt"
)

func main() {
	config.Init()

	database.Init()
	defer database.Close()

	m := database.GetM()

	err := m.Down()
	if err != nil {
		panic(err)
	}

	fmt.Println("Executed m.Down() !!")
}
