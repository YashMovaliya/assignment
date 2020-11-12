package main

import (
	"github.com/YashMovaliya/assignment/go-todo-restapi/app"
	"github.com/YashMovaliya/assignment/go-todo-restapi/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
