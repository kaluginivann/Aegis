package main

import (
	"github.com/kaluginivann/Aegis/internal/configs"
	e "github.com/kaluginivann/Aegis/internal/engine"
)

func main() {
	config := configs.LoadConfig()
	engine := e.NewEngine(config)
	engine.Run()
}
