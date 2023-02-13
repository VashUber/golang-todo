package main

import (
	"github.com/VashUber/golang-todo/infrastructure"
	"github.com/VashUber/golang-todo/internal/app"
)

func main() {
	cfg := infrastructure.NewConfig()

	app.Run(cfg)
}
